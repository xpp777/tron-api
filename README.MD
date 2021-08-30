# tron api sdk

## 地址

### 生成地址 私钥

```
 addr, PrivateKey, err := address.CreatAddress(pwd)
	 if err != nil {
	 	fmt.Printf("err :%s", err.Error())
	 	return
	 }
```

### 解密私钥

```
    PrivateKey, err := address.GetPrivateKey(pwd, PrivateKey)
     if err != nil {
	 	fmt.Printf("err :%s", err.Error())
	 	return
	 }
```

### 验证地址

```
    if address.ValidAddress(addr){
        fmt.Println("地址：合法")
    }else{
        fmt.Println("地址不合法")
    }

```

## 获取指定区块取件

```
    ApiKeys := []string{
		"9e3b28c4-3fd3-48c9-97af-ce4af055bcbb",
		"e29b2bc3-acba-4fe1-b784-ea54dafd0b6d",
	}
	client := api.NewClient(33251080,"",ApiKeys)
	client.GetBlockByLimitNext(func(data *model.TransferData) {
		fmt.Println(data)
	})
    select {}
```

## 发起交易

```
    // 初始化设置
   ContractMap := make(map[string]*model.ContractModel)
	ContractMap["TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"] = &model.ContractModel{
		Name:     "USDT",
		Type:     "trc20",
		Contract: "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
		Decimal:  6,
	}
	rpc.SetContractMap(ContractMap)
    // 发起交易
	txid, err := rpc.NewRpc().Sen(PrivateKey, "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t", "TD1HSXrv756gehQj1fi9hEUfVkheR9G4DV", decimal.NewFromFloat(0.01))
	if err != nil {
		g.Log().Error(err)
		return
	}
	fmt.Println(txid)
```