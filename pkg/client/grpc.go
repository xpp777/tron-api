package client

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/shopspring/decimal"
	"github.com/xiaomingping/tron-api/pkg/base58"
	"github.com/xiaomingping/tron-api/pkg/crypto"
	"github.com/xiaomingping/tron-api/pkg/hexutil"
	"google.golang.org/grpc"
	"log"
	"math/big"
	"sync"
	"time"
)

var (
	feelimit int64 = 40000000 // 转账合约燃烧 trx数量 单位 sun 默认0.5trx 转账一笔大概消耗能量 0.26trx
	Trx            = "trx"
	Urls           = []string{
		"3.225.171.164",
		"52.53.189.99",
		"18.196.99.16",
		"34.253.187.192",
		"18.133.82.227",
		"35.180.51.163",
		"54.252.224.209",
		"18.228.15.36",
		"52.15.93.92",
		"34.220.77.106",
		"13.127.47.162",
		"13.124.62.58",
		"35.182.229.162",
		"18.209.42.127",
		"3.218.137.187",
		"34.237.210.82",
	}
	connIndex    int
	connMutex    sync.Mutex
)

type Rpc struct {
	Client api.WalletClient
	Conn   *grpc.ClientConn
}

// 获取超时上下文
func (r *Rpc) timeoutContext() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	go func() {
		time.Sleep(time.Second * 30)
		cancel()
	}()
	return ctx
}

// 轮询节点
func getPollingUrl() string {
	connMutex.Lock()
	defer connMutex.Unlock()
	lens := len(Urls)
	if connIndex >= lens {
		connIndex = 0
	}
	inst := Urls[connIndex]
	connIndex = (connIndex + 1) % lens
	return inst + ":50051"
}

// 获取新节点
func GetNode() (*Rpc, error) {
	Conn, err := grpc.Dial(getPollingUrl(), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	Client := api.NewWalletClient(Conn)
	return &Rpc{Conn: Conn, Client: Client}, nil
}
// 处理合约获取余额参数
func (r *Rpc) processBalanceOfParameter(addr string) (data []byte) {
	methodID, _ := hexutil.Decode("70a08231")
	add, _ := base58.DecodeCheck(addr)
	paddedAddress := common.LeftPadBytes(add[1:], 32)
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	return
}

// 处理合约获取余额
func (r *Rpc) processBalanceOfData(trc20 []byte) (amount int64) {
	if len(trc20) >= 32 {
		amount = new(big.Int).SetBytes(common.TrimLeftZeroes(trc20[0:32])).Int64()
	}
	return
}

// 处理合约参数
func (r *Rpc) processTransferData(trc20 []byte) (to string, amount int64, flag bool) {
	if len(trc20) >= 68 {
		if hexutil.Encode(trc20[:4]) != "a9059cbb" {
			return
		}
		// 多1位41
		trc20[15] = 65
		to = base58.EncodeCheck(trc20[15:36])
		amount = new(big.Int).SetBytes(common.TrimLeftZeroes(trc20[36:68])).Int64()
		flag = true
	}
	return
}

// 处理合约转账参数
func (r *Rpc) processTransferParameter(to string, amount int64) (data []byte) {
	methodID, _ := hexutil.Decode("a9059cbb")
	addr, _ := base58.DecodeCheck(to)
	paddedAddress := common.LeftPadBytes(addr[1:], 32)
	amountBig := new(big.Int).SetInt64(amount)
	paddedAmount := common.LeftPadBytes(amountBig.Bytes(), 32)
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	return
}

// trx转账
func (r *Rpc) transfer(ownerKey *ecdsa.PrivateKey, toAddress string, amount int64) (string, error) {
	transferContract := new(core.TransferContract)
	transferContract.OwnerAddress = crypto.PubkeyToAddress(ownerKey.
		PublicKey).Bytes()
	transferContract.ToAddress, _ = base58.DecodeCheck(toAddress)
	transferContract.Amount = amount

	transferTransactionEx, err := r.Client.CreateTransaction2(r.timeoutContext(), transferContract)

	var txid string
	if err != nil {
		return txid, err
	}
	transferTransaction := transferTransactionEx.Transaction
	if transferTransaction == nil || len(transferTransaction.
		GetRawData().GetContract()) == 0 {
		return txid, fmt.Errorf("transfer error: invalid transaction")
	}
	hash, err := SignTransaction(transferTransaction, ownerKey)
	if err != nil {
		return txid, err
	}
	txid = hexutil.Encode(hash)

	result, err := r.Client.BroadcastTransaction(r.timeoutContext(),
		transferTransaction)
	if err != nil {
		return "", err
	}
	if !result.Result {
		return "", fmt.Errorf("api get false the msg: %v", result.String())
	}
	return txid, err
}

// 合约转账 TRC20
func (r *Rpc) transferContract(ownerKey *ecdsa.PrivateKey, Contract string, data []byte, feeLimit int64) (string, error) {
	transferContract := new(core.TriggerSmartContract)
	transferContract.OwnerAddress = crypto.PubkeyToAddress(ownerKey.
		PublicKey).Bytes()
	transferContract.ContractAddress, _ = base58.DecodeCheck(Contract)
	transferContract.Data = data
	transferTransactionEx, err := r.Client.TriggerConstantContract(r.timeoutContext(), transferContract)
	var txid string
	if err != nil {
		return txid, err
	}
	transferTransaction := transferTransactionEx.Transaction
	if transferTransaction == nil || len(transferTransaction.
		GetRawData().GetContract()) == 0 {
		return txid, fmt.Errorf("transfer error: invalid transaction")
	}
	if feeLimit > 0 {
		transferTransaction.RawData.FeeLimit = feeLimit
	}

	hash, err := SignTransaction(transferTransaction, ownerKey)
	if err != nil {
		return txid, err
	}
	txid = hexutil.Encode(hash)

	result, err := r.Client.BroadcastTransaction(r.timeoutContext(),
		transferTransaction)
	if err != nil {
		return "", err
	}
	if !result.Result {
		return "", fmt.Errorf("api get false the msg: %v", result.String())
	}
	return txid, err
}

// 转账
func (r *Rpc) Sen(key *ecdsa.PrivateKey, contract, to string, amount decimal.Decimal) (string, error) {
	Type, Decimal := ChargeContract(contract)
	switch Type {
	case Trx:
		var amountdecimal = decimal.New(1, Decimal)
		amountac, _ := amount.Mul(amountdecimal).Float64()
		return r.transfer(key, to, int64(amountac))
	case "trc20":
		var amountdecimal = decimal.New(1, Decimal)
		amountac, _ := amount.Mul(amountdecimal).Float64()
		data := r.processTransferParameter(to, int64(amountac))
		return r.transferContract(key, contract, data, feelimit)
	case "trc10":
		return "", nil
	default:
		return "", nil
	}
}
