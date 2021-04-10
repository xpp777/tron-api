package client

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/golang/protobuf/proto"
	"time"
	"tron-api/pkg/crypto"
)

// SignTransaction 签名交易
func SignTransaction(transaction *core.Transaction, key *ecdsa.PrivateKey) ([]byte, error) {
	transaction.GetRawData().Timestamp = time.Now().UnixNano() / 1000000
	rawData, err := proto.Marshal(transaction.GetRawData())
	if err != nil {
		return nil, err
	}
	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)
	contractList := transaction.GetRawData().GetContract()
	for range contractList {
		signature, err := crypto.Sign(hash, key)
		if err != nil {
			return nil, err
		}
		transaction.Signature = append(transaction.Signature, signature)
	}
	return hash, nil
}