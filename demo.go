package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xiaomingping/tron-api/pkg/address"
	"github.com/xiaomingping/tron-api/pkg/client/api"
)
const PrivateKey = "wZUfIRm5J3Bxeivt"


func main() {
	Asset := "hDwDjgW5p8XQhLImeD1D2E4HcO8nNsAUe0JM9NLw/wmQbTBoQMDxy17Y6u7TsigKdsnRHjJMI9jtdv6mFRcjND88mJlyTIJjDAeuL2YN8/M="
	PrivateKeys, err := address.GetPrivateKey(PrivateKey, Asset)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ClientApi := api.NewClient(0,"",[]string{
		"9e3b28c4-3fd3-48c9-97af-ce4af055bcbb",
		"e29b2bc3-acba-4fe1-b784-ea54dafd0b6d",
		"259625c9-112b-4406-a4f5-bd392d44f90b",
		"e07ae70a-5925-41fd-84c4-e8a3757cf6e7",
		"9e3b28c4-3fd3-48c9-97af-ce4af055bcbb",
		"e29b2bc3-acba-4fe1-b784-ea54dafd0b6d",
		"259625c9-112b-4406-a4f5-bd392d44f90b",
	})
	val := decimal.NewFromFloat(0.1)
	res,err := ClientApi.Sen(PrivateKeys,"trx","TAgShRz67UnfwuMk1UHKgg9PkEWe3r8igG",val)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
}