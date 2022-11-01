package types

import "time"

type SubgraphToken struct {
	Address   string `json:"address"`   // The address of the token
	Balance   string `json:"balance"`   // The balance of the token in the pool
	Decimals  int    `json:"decimals"`  // The number of decimals of the token
	PriceRate string `json:"priceRate"` // The price rate of the token in the pool
	Weight    string `json:"weight"`    // The weight of the token in the pool
}

type SubgraphPoolBase struct {
	Id          string          `json:"id"`          // pool id
	Address     string          `json:"address"`     // pool address
	PoolType    string          `json:"poolType"`    // pool type
	SwapFee     string          `json:"swapFee"`     // pool swap fee
	SwapEnabled bool            `json:"swapEnabled"` // pool swap enabled
	TotalShares string          `json:"totalShares"` // pool total shares
	Tokens      []SubgraphToken `json:"tokens"`      // pool tokens
	TokensList  []string        `json:"tokensList"`  // pool tokens list
	// Weighted & Element field
	TotalWeight string `json:"totalWeight"` // pool total weight
	// Stable specific fields
	Amp string `json:"amp"` // pool amp
	// Element specific fields
	ExpiryTime     time.Duration `json:"expiryTime"`     // pool expiry time
	UnitSeconds    int           `json:"unitSeconds"`    // pool unit seconds
	PrincipalToken string        `json:"principalToken"` // pool principal token
	BaseToken      string        `json:"baseToken"`      // pool base token
	// Linear specific fields
	MainIndex    int    `json:"mainIndex"`    // pool main index
	WrappedIndex int    `json:"wrappedIndex"` // pool wrapped index
	LowerTarget  string `json:"lowerTarget"`  // pool lower target
	UpperTarget  string `json:"upperTarget"`  // pool upper target
	// Gyro2 specific field
	//gyro2PriceBounds Gyro2PriceBounds
	// Gyro3 specific field
	//gyro3PriceBounds Gyro3PriceBounds
}
