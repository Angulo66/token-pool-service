package math

type PoolToken struct {
	Name     string `json:"name"`
	Address  string `json:"id"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	Balance  string `json:"balance"`
	Weight   string `json:"weight"`
}

type Pool struct {
	Id                string    `json:"id"`
	Token0            PoolToken `json:"token0"`
	Token1            PoolToken `json:"token1"`
	TrackedReserveETH string    `json:"trackedReserveETH"`
	ReservedETH       string    `json:"reservedETH"`
	ReserveUSD        string    `json:"reserveUSD"`
	TotalSupply       string    `json:"totalSupply"`
	Reserve0          string    `json:"reserve0"`
	Reserve1          string    `json:"reserve1"`
}
