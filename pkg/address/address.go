package address

import (
	"crypto/ecdsa"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/xiaomingping/tron-api/pkg/base58"
	"github.com/xiaomingping/tron-api/pkg/crypto"
	"github.com/xiaomingping/tron-api/pkg/hexutil"
	"github.com/xiaomingping/tron-api/pkg/keystore"
	"strings"
)

// 生成地址 私钥
func CreatAddress(pwd string) (addr, PrivateKey string, err error) {
	re, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}
	addr = base58.EncodeCheck(crypto.PubkeyToAddress(re.PublicKey).Bytes())
	password := keystore.HashAndSalt([]byte(pwd + "trx"))
	prikey := crypto.PrikeyToHexString(re)
	md5sum := md5.Sum([]byte(password))
	result, err1 := crypto.AesEncrypt([]byte(prikey), md5sum[:])
	if err1 != nil {
		err = err1
		return
	}
	return addr, base64.StdEncoding.EncodeToString(result), err
}

func Encrypt(pwd, PrivateKey string) (string, error) {
	password := keystore.HashAndSalt([]byte(pwd + "trx"))
	md5sum := md5.Sum([]byte(password))
	result, err := crypto.AesEncrypt([]byte(PrivateKey), md5sum[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(result), err
}

func GetPrivateKey(pwd, PrivateKey string) (account *ecdsa.PrivateKey, err error) {
	password := keystore.HashAndSalt([]byte(pwd + "trx"))
	re, err1 := base64.StdEncoding.DecodeString(PrivateKey)
	if err1 != nil {
		err = err1
		return
	}
	md5sum := md5.Sum([]byte(password))
	result, err1 := crypto.AesDecrypt(re, md5sum[:])
	if err1 != nil {
		err = err1
		return
	}
	fmt.Println(string(result))
	account, err = crypto.GetPrivateKeyByHexString(string(result))
	return
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

func HexTOString(HexAddress string) string {
	add := strings.Replace(HexAddress, "0x", "41", 1)
	hex, _ := hexutil.Hex2Bytes(add)
	return base58.EncodeCheck(hex)
}
