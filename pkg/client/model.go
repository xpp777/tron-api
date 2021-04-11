package client

import "github.com/shopspring/decimal"

type BalanceModel struct {
	Trx  string `json:"trx"`  // trx 余额
	GLV  string `json:"glv"`  // glv 余额
	Usdt string `json:"usdt"` // u 余额
}

// Contract 合约 TRC20 和 TRC10
type Contract struct {
	Name                string          `toml:"name"`       // USDT BTT
	Type                string          `toml:"type"`       // TRC20 和 TRC10
	Contract            string          `toml:"contract"`   // 合约地址或者合约ID
	Decimal             int32           `toml:"decimal"`    // 合约小数位
	CollectionMinAmount decimal.Decimal `toml:"min_amount"` // 代币最小归集数目
}

// 用户信息
type GetAccountModel struct {
	LatestOprationTime int64 `json:"latest_opration_time"`
	OwnerPermission    struct {
		Keys           []KeysAddress `json:"keys"`
		Threshold      int           `json:"threshold"`
		PermissionName string        `json:"permission_name"`
	} `json:"owner_permission"`
	FreeAssetNetUsagev2 []Keys `json:"free_asset_net_usageV2"`
	AccountResource     struct {
		LatestConsumeTimeForEnergy int64 `json:"latest_consume_time_for_energy"`
	} `json:"account_resource"`
	ActivePermission []struct {
		Operations     string        `json:"operations"`
		Keys           []KeysAddress `json:"keys"`
		Threshold      int           `json:"threshold"`
		ID             int           `json:"id"`
		Type           string        `json:"type"`
		PermissionName string        `json:"permission_name"`
	} `json:"active_permission"`
	Assetv2               []Keys              `json:"assetV2"`
	Address               string              `json:"address"`
	Balance               int                 `json:"balance"`
	CreateTime            int64               `json:"create_time"`
	Trc20                 []map[string]string `json:"trc20"`
	LatestConsumeFreeTime int64               `json:"latest_consume_free_time"`
}

type RespAccount struct {
	Data    []GetAccountModel `json:"data"`
	Success bool              `json:"success"`
	Meta    Meta              `json:"meta"`
}

type Meta struct {
	At       int64 `json:"at"`
	PageSize int   `json:"page_size"`
}

type Keys struct {
	Value int    `json:"value"`
	Key   string `json:"key"`
}
type KeysAddress struct {
	Address string `json:"address"`
	Weight  int    `json:"weight"`
}

type RespTransactionsTrc20 struct {
	Data    []TransactionsTrc20Model `json:"data"`
	Success bool                     `json:"success"`
	Meta    struct {
		At       int64 `json:"at"`
		PageSize int   `json:"page_size"`
	} `json:"meta"`
}

type TransactionsTrc20Model struct {
	TransactionID  string    `json:"transaction_id"`
	TokenInfo      TokenInfo `json:"token_info"`
	BlockTimestamp int64     `json:"block_timestamp"`
	From           string    `json:"from"`
	To             string    `json:"to"`
	Type           string    `json:"type"`
	Value          string    `json:"value"`
}

type TokenInfo struct {
	Symbol   string `json:"symbol"`
	Address  string `json:"address"`
	Decimals int32  `json:"decimals"`
	Name     string `json:"name"`
}

type GettransactioninfobyidModel struct {
	ID              string   `json:"id"`
	Fee             int      `json:"fee"`
	Blocknumber     int      `json:"blockNumber"`
	Blocktimestamp  int64    `json:"blockTimeStamp"`
	Contractresult  []string `json:"contractResult"`
	ContractAddress string   `json:"contract_address"`
	Receipt         struct {
		EnergyUsage       int    `json:"energy_usage"`
		EnergyFee         int    `json:"energy_fee"`
		OriginEnergyUsage int    `json:"origin_energy_usage"`
		EnergyUsageTotal  int    `json:"energy_usage_total"`
		NetFee            int    `json:"net_fee"`
		Result            string `json:"result"`
	} `json:"receipt"`
	Log []struct {
		Address string   `json:"address"`
		Topics  []string `json:"topics"`
		Data    string   `json:"data"`
	} `json:"log"`
}


// Contract 合约 TRC20 和 TRC10
type ContractModel struct {
	Name                string          `toml:"name"`       // USDT BTT
	Type                string          `toml:"type"`       // TRC20 和 TRC10
	Contract            string          `toml:"contract"`   // 合约地址或者合约ID
	Decimal             int32           `toml:"decimal"`    // 合约小数位
}
