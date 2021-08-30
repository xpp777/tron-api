package main

func main() {
	// client.Urls = []string{
	// 	"grpc.shasta.trongrid.io",
	// 	"grpc.shasta.trongrid.io",
	// 	"grpc.shasta.trongrid.io",
	// }
	// path := "./key"
	// pwd := "a"
	// addr, PrivateKey, err := address.CreatAddress(pwd)
	// if err != nil {
	// 	fmt.Printf("err :%s", err.Error())
	// 	return
	// }
	// fmt.Printf("addr:%s,PrivateKey:%s\n", addr, PrivateKey)
	// PrivateKey := ""
	// res,err := address.GetPrivateKey(pwd,PrivateKey)
	// if err != nil {
	// 	fmt.Printf("err :%s", err.Error())
	// 	return
	// }
	// addr := base58.EncodeCheck(crypto.PubkeyToAddress(res.PublicKey).Bytes())
	// fmt.Println(addr)

	//ci,err := client.GetNode()
	//if err != nil {
	//		fmt.Printf("err :%s", err.Error())
	//	return
	//}
	//ci.GetNowBlock2(func(data *client.TransferData) {
	//	fmt.Println(data)
	//})
	// re,err := ci.GetTrc20Balance("TLYUprahhotHaKQ9U4s3AiXh7S5vuiMtBi",addr,6,res)
	// if err != nil {
	// 	fmt.Printf("err :%s", err.Error())
	// 	return
	// }
	// client.ApiKeys = []string{
	// 	"5527c743-dc35-4a00-8b97-7e75ac9c164b",
	// 	"4c492539-5e03-452b-9633-6e5b8998cc36",
	// }
	// var Contract []client.ContractModel
	// Contract = append(Contract, client.ContractModel{
	// 	Name:               "glv",
	// 	Type:                "trc20",
	// 	Contract:            "TLYUprahhotHaKQ9U4s3AiXh7S5vuiMtBi",
	// 	Decimal:             6,
	// })
	// client.InitContract(Contract)
	// c := client.Client{Url: client.ApiUrlShasta}
	// c := client.NewClient()
	//  resp,err := c.GetAccount("TJgMDURAhPEcDw8JXpJKttSSK6f3zrYgED")
	// if err != nil {
	// 	fmt.Printf("err :%s", err.Error())
	// 	return
	// }
	// fmt.Println(resp.Balance)
	// fmt.Println(client.GetTRXBalance(resp))
	// client.ApiKeys = []string{
	// 	"5527c743-dc35-4a00-8b97-7e75ac9c164b",
	// 	"4c492539-5e03-452b-9633-6e5b8998cc36",
	// }
	// Resp, err := c.GetTransactionsTrc20("TD3y5r2AxHfdjBXA336GVVqxYrjNMKgZrn", "TLYUprahhotHaKQ9U4s3AiXh7S5vuiMtBi")
	// if err != nil {
	// 	fmt.Printf("err :%s", err.Error())
	// 	return
	// }
	// resp,err := c.GetBlockById("")
	// if err != nil {
	// 	fmt.Printf("err :%s", err.Error())
	// 	return
	// }
	// fmt.Println(Resp)
	// for _, v := range Resp {
	// 	fmt.Println(v.TransactionID)
	// 	fmt.Println(client.BalanceAccuracy(v.Value, -v.TokenInfo.Decimals))
	// }
	// resp, err := c.GetAccount("TD3y5r2AxHfdjBXA336GVVqxYrjNMKgZrn")
	// if err != nil {
	// 	fmt.Printf("err :%s", err.Error())
	// 	return
	// }
	// fmt.Println(resp.Balance)
	// fmt.Println(client.GetTRXBalance(resp))
	// PrivateKey,err := crypto.HexToECDSA("")
	// if err != nil {
	// 	fmt.Errorf("err:%v", err)
	// 	return
	// }
	// val := decimal.NewFromFloat(10.1)
	// resp,err := ci.Sen(PrivateKey,client.Trx,"TDJxbBwXpmPMBvGjNPjRH1urvhuCWeETgq",val)
	// if err != nil {
	// 		fmt.Printf("err :%s", err.Error())
	// }
	// ci,err := client.GetNode()
	// if err != nil {
	// 		fmt.Printf("err :%s", err.Error())
	// }
	// // resp,err := ci.GetAccount("TX17ncfbkyTMMxHWaiGBqDuGJP987WaEuE")
	// // if err != nil {
	// // 	fmt.Printf("err :%s", err.Error())
	// // }
	// // fmt.Println(resp.OwnerPermission)
	// // fmt.Println(resp.WitnessPermission)
	// // fmt.Println(resp.ActivePermission)
	// resp,err := ci.GetBlockById("")
	// if err != nil {
	// 	fmt.Printf("err :%s", err.Error())
	// }
	// fmt.Println(resp.RawData.RefBlockNum)
	// fmt.Println(resp.Ret[0])
	// client.Urls = []string{
	// 	"grpc.shasta.trongrid.io",
	// 	"grpc.shasta.trongrid.io",
	// 	"grpc.shasta.trongrid.io",
	// }
	// ci,err := client.GetNode()
	// if err != nil {
	// 		fmt.Printf("err :%s", err.Error())
	// }
	// client.SetStartNum(14428235)
	// for  {
	// 	fmt.Println(client.GetStartNum())
	// 	ci.GetNowBlock2(func(data *client.TransferData) {
	// 		if data.Contract == "trx" {
	// 			return
	// 		}
	// 		fmt.Println(data)
	// 	})
	// 	time.Sleep(time.Second * 20)
	// }
	// resp,err := ci.GetBlockById("4333d42a00ab2f5fdabfdef5b4d98df81dc256d44392735fc4e0f5a57677f663")
	// if err != nil {
	// 	fmt.Printf("err :%s", err.Error())
	// 	return
	// }
	// fmt.Println(resp.GetExchangeId())
	// fmt.Println(resp.GetBlockTimeStamp())
	// fmt.Println(resp.GetContractAddress())
	// fmt.Println(resp.GetExchangeReceivedAmount())
	// resps, err := ci.GetAccount("TX17ncfbkyTMMxHWaiGBqDuGJP987WaEuE")
	// if err != nil {
	// 	fmt.Printf("err :%s", err.Error())
	// 	return
	// }
	// fmt.Println(resps.Balance)
	// fmt.Println(resps.GetAsset())

	// 地址转换
	// add := strings.Replace("0x0970b863a1cb850e5bdb27e8653057abfe94aadc", "0x", "41", 1)
	// hex,_ := hexutil.Hex2Bytes(add)
	// add = base58.EncodeCheck(hex)
	// fmt.Println("add",add)
	//
	//
	//var Contract []client.ContractModel
	//Contract = append(Contract, client.ContractModel{
	//	Name:     "USDT",
	//	Type:     "trc20",
	//	Contract: "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
	//	Decimal:  6,
	//})
	//err := client.InitContract(Contract)
	//if err != nil {
	//	fmt.Printf("err :%s", err.Error())
	//	return
	//}
	//client.ApiKeys = []string{
	//	"9e3b28c4-3fd3-48c9-97af-ce4af055bcbb",
	//	"e29b2bc3-acba-4fe1-b784-ea54dafd0b6d",
	//}
	//Client := client.NewClient()
	// // TD91gHfn4xML3LcuiZsiD1Q7wcDBxJiiFv
	// Account,err := Client.GetAccount("TGx7FL5fT2BZi1obiuCNdq7yUd22LtbMTU")
	// if err != nil {
	// 	fmt.Printf("err :%s", err.Error())
	// 	return
	// }
	// data := client.GetTRXBalance(Account)
	// fmt.Println(data)
	// res,err := Client.GetBlockById("709c6ad9ee72d455c533b7a379f2cd8dac49df37d1774f507428801bf67964c6")
	// if err != nil {
	// 	fmt.Printf("err :%s", err.Error())
	// 	return
	// }
	// msg,_ := json.Marshal(res)
	// fmt.Println(string(msg))
	//data := processTransferParameter("41b6011a31721cd1a73c79c8e747ff5882b84a5376",2000000000)
	//g.Log().Println(string(data))
	//ApiKeys := []string{
	//	"9e3b28c4-3fd3-48c9-97af-ce4af055bcbb",
	//	"e29b2bc3-acba-4fe1-b784-ea54dafd0b6d",
	//}
	//client := api.NewClient(	33251080,"",ApiKeys)
	//client.GetBlockByLimitNext(func(data *model.TransferData) {
	//	fmt.Println(data)
	//})
	//select {}
}