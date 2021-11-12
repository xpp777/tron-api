package api

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/shopspring/decimal"
	"github.com/xiaomingping/tron-api/pkg/address"
	"github.com/xiaomingping/tron-api/pkg/model"
	"github.com/xiaomingping/tron-api/pkg/sign"
	"math/big"
	"sync"
)

const (
	ApiUrl       = "https://api.trongrid.io"        // 主网
	ApiUrlShasta = "https://api.shasta.trongrid.io" // Shasta测试网
)

var (
	curIndex        = 0
	mutex           sync.Mutex
	Trx                   = "trx"
	trxDecimal      int32 = 6 // trx 单位
	mapContract           = make(map[string]*model.ContractModel)
	mapContractType       = map[string]bool{
		"trx":   true,
		"trc10": true,
		"trc20": true,
	}
)

type Client struct {
	Url      string // 请求地址
	mutex    sync.Mutex
	count    int      // 每次获取块数量
	apiKeys  []string // api key
	startNum int      // 起始块
}

func SetContractMap(ContractMap map[string]*model.ContractModel) {
	mapContract = ContractMap
}

// 判断当前属于什么合约
func ChargeContract(contract string) (string, int32) {
	if contract == "trx" || contract == "" {
		return Trx, trxDecimal
	}
	if v := mapContract[contract]; v != nil {
		if ok, _ := mapContractType[v.Type]; ok {
			return v.Type, v.Decimal
		}
	}
	return "NONE", 18
}
func NewClient(Num int, apiUrl string, apiKeys []string) *Client {
	if apiUrl == "" {
		apiUrl = ApiUrl
	}
	return &Client{
		Url:      apiUrl,
		count:    20,
		startNum: Num,
		apiKeys:  apiKeys,
	}
}

// 获取api key
func (c *Client) getApiKey() string {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	lens := len(c.apiKeys)
	if curIndex >= lens {
		curIndex = 0
	}
	inst := c.apiKeys[curIndex]
	curIndex = (curIndex + 1) % lens
	return inst
}

// 获取请求客户端
func (c *Client) getClient() *ghttp.Client {
	Client := ghttp.NewClient()
	Client.SetHeader("Content-Type", "application/json")
	Client.SetHeader("TRON-PRO-API-KEY", c.getApiKey())
	return Client
}
func (c *Client) SetStartNum(StartNum int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.startNum = StartNum
}
func (c *Client) GetStartNum() int {
	return c.startNum
}

// 获取最新块数据
func (c *Client) GetBlockByLimitNext(Transfer func(*model.TransferData)) {
	URL := fmt.Sprintf("%s/wallet/getblockbylimitnext", c.Url)
	body := c.getClient().PostVar(URL, map[string]interface{}{"startNum": c.startNum, "endNum": c.startNum + c.count, "visible": true})
	if body.IsEmpty() {
		return
	}
	var NewBlock model.NewBlock
	err := body.Structs(&NewBlock)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	StartNum := c.startNum + len(NewBlock.Block)
	c.SetStartNum(StartNum)
	for _, v := range NewBlock.Block {
		c.processBlock(v, Transfer)
	}
}

// 合约解析
func (c *Client) processBlock(block model.Block, Transfer func(*model.TransferData)) {
	for _, v := range block.Transactions {
		txId := v.TxID
		if v.Ret == nil || len(v.Ret) == 0 {
			continue
		}
		rets := v.Ret[0].ContractRet
		if rets != "SUCCESS" {
			continue
		}
		for _, val := range v.RawData.Contract {
			switch val.Type {
			case "TransferContract": // trx 转账
				unObj := &model.Trc{}
				err := gconv.Struct(val.Parameter.Value, unObj)
				if err != nil {
					fmt.Printf("parse Contract %v err: %v\n", val, err)
					continue
				}
				Transfer(&model.TransferData{FormAddress: unObj.OwnerAddress, ToAddress: unObj.ToAddress, Amount: unObj.Amount, Contract: "trx", TxId: txId})
			case "TransferAssetContract": // trc10
				continue
			case "TriggerSmartContract":
				// trc20 转账
				unObj := &model.Value{}
				err := gconv.Struct(val.Parameter.Value, unObj)
				if err != nil {
					fmt.Printf("parse Contract %v err: %v\n", val, err)
					continue
				}
				to, amount, flag := c.processTransferData(unObj.Data)
				if flag { // 只有调用了 transfer(address,uint256) 才是转账
					Transfer(&model.TransferData{FormAddress: unObj.OwnerAddress, ToAddress: to, Amount: amount, Contract: unObj.ContractAddress, TxId: txId})
				}
			}
		}
	}
}

// 处理合约参数
func (c *Client) processTransferData(trc20 string) (to string, amount int64, flag bool) {
	if len(trc20) >= 136 {
		if trc20[:8] != "a9059cbb" {
			return
		}
		addressHex := trc20[32:72]
		if len(addressHex) != 42 {
			addressHex = "41" + addressHex
		}
		to = address.HexTOString(addressHex)
		s := string(sign.TrimLeftZeroes([]byte(trc20[72:])))
		var b *big.Int
		b, flag = new(big.Int).SetString(s, 16)
		if flag {
			amount = b.Int64()
		}
	}
	return
}

// 进度转换
func (c *Client) BalanceAccuracy(Balance string, exp int32) string {
	b, _ := decimal.NewFromString(Balance)
	return b.Mul(decimal.New(1, exp)).String()
}

// 获取用户信息
func (c *Client) GetAccount(address string) (map[string]string, error) {
	url := fmt.Sprintf("%s/v1/accounts/%s", c.Url, address)
	body := c.getClient().GetVar(url)
	if body.IsEmpty() {
		return nil, errors.New("网络错误")
	}
	var Account model.RespAccount
	err := body.Struct(&Account)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if Account.Success != true {
		return nil, errors.New("连接失败")
	}
	if len(Account.Data) == 0 {
		return nil, errors.New("账号未激活")
	}
	data := Account.Data[0]
	BalanceModel := make(map[string]string)
	BalanceModel["trx"] = gconv.String(data.Balance)
	for _, Tokens := range data.Trc20 {
		for key, val := range Tokens {
			BalanceModel[key] = val
		}
	}
	return BalanceModel, nil
}

// 获取账户历史TRC20交易记录
func (c *Client) GetTransactionsTrc20(address, contract string) ([]model.TransactionsTrc20Model, error) {
	url := fmt.Sprintf("%s/v1/accounts/%s/transactions/trc20?only_confirmed=true&only_to=true&contract_address=%s", c.Url, address, contract)
	body := c.getClient().GetVar(url)
	if body.IsEmpty() {
		return nil, errors.New("网络错误")
	}
	var TransactionsTrc20 model.RespTransactionsTrc20
	err := body.Struct(&TransactionsTrc20)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if TransactionsTrc20.Success != true {
		return nil, errors.New("连接失败")
	}
	return TransactionsTrc20.Data, nil
}

// 获取区块详情
func (c *Client) GetBlockById(exchangeId string) (*model.ContractBlockInfo, error) {
	url := fmt.Sprintf("%s/event/transaction/%s", c.Url, exchangeId)
	body := c.getClient().GetVar(url)
	if body.IsEmpty() {
		return nil, errors.New("网络错误")
	}
	var AutoGenerated []model.AutoGenerated
	err := body.Structs(&AutoGenerated)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if len(AutoGenerated) == 0 {
		return nil, errors.New("没有此交易")
	}
	return &model.ContractBlockInfo{
		TransactionID:   AutoGenerated[0].TransactionID,
		BlockNumber:     AutoGenerated[0].BlockNumber,
		EventName:       AutoGenerated[0].EventName,
		BlockTimestamp:  AutoGenerated[0].BlockTimestamp,
		ContractAddress: AutoGenerated[0].ContractAddress,
		From:            address.HexTOString(AutoGenerated[0].Result.From),
		To:              address.HexTOString(AutoGenerated[0].Result.To),
		Value:           AutoGenerated[0].Result.Value,
		EventIndex:      AutoGenerated[0].EventIndex,
	}, nil
}

