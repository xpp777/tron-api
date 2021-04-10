package address

import (
	"github.com/xiaomingping/tron-api/pkg/base58"
	"github.com/xiaomingping/tron-api/pkg/crypto"
	"github.com/xiaomingping/tron-api/pkg/keystore"
)

// 生成地址 私钥
func CreatAddress(path, pwd string) (addr, filePath string, err error) {
	re, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}
	addr = base58.EncodeCheck(crypto.PubkeyToAddress(re.PublicKey).Bytes())
	paths, err := keystore.StoreAccountToKeyStoreFile(re, pwd, path+"/"+addr)
	if err != nil {
		return "", "", err
	}
	return addr, paths, err
}

// 验证地址
func ValidAddress(addr string) bool {
	if len(addr) != 34 {
		return false
	}
	if string(addr[0:1]) != "T" {
		return false
	}
	_, err := base58.DecodeCheck(addr)
	if err != nil {
		return false
	}
	return true
}
