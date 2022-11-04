package math

import (
	"math"
	"strconv"
	"strings"

	etypes "github.com/angulo66/subgraph-pool-service/pkg/balancer/types"
)

type WeightedPoolPairData struct {
	BalanceIn   float64 `json:"balanceIn"`
	BalanceOut  float64 `json:"balanceOut"`
	WeightIn    float64 `json:"weightIn"`
	WeightOut   float64 `json:"weightOut"`
	SwapFee     float64 `json:"swapFee"`
	DecimalsIn  int     `json:"decimalsIn"`
	DecimalsOut int     `json:"decimalsOut"`
}

func GetSwapData(amount float64, subgraphPool *etypes.SubgraphPoolBase, tokenIn string, tokenOut string) []float64 {
	wp := PoolDataToWeightedPoolPairData(subgraphPool, tokenIn, tokenOut)
	ep := ExactTokenInForTokenOut(amount, wp)
	sp := GetSpotPrice(amount, wp)
	eP := GetEffectivePrice(amount, wp)
	pi := GetPriceImpact(eP, sp)
	return []float64{ep, sp, eP, pi}
}

// https://github.com/balancer-labs/balancer-sor/blob/john/v2-package-linear/src/pools/weightedPool/weightedMath.ts
func ExactTokenInForTokenOut(amount float64, poolPairData *WeightedPoolPairData) float64 {
	Bi := poolPairData.BalanceIn
	Bo := poolPairData.BalanceOut
	wi := poolPairData.WeightIn
	wo := poolPairData.WeightOut
	Ai := amount
	f := poolPairData.SwapFee

	return Bo * (1 - math.Pow(Bi/(Bi+Ai*(1-f)), (wi/wo)))
}

func GetEffectivePrice(amount float64, poolPairData *WeightedPoolPairData) float64 {
	Bi := poolPairData.BalanceIn
	Bo := poolPairData.BalanceOut
	wi := poolPairData.WeightIn
	wo := poolPairData.WeightOut
	Ai := amount
	f := poolPairData.SwapFee

	return Ai * (1 - f) / (Bo * (1 - math.Pow(Bi/(Bi+Ai*(1-f)), (wi/wo))))
}

func GetSpotPrice(amount float64, poolPairData *WeightedPoolPairData) float64 {
	Bi := poolPairData.BalanceIn
	Bo := poolPairData.BalanceOut
	wi := poolPairData.WeightIn
	wo := poolPairData.WeightOut

	return (Bi / wi) / (Bo / wo)
}

func GetPriceImpact(ep float64, sp float64) float64 {
	return 1 - (ep / sp)
}

func PoolDataToWeightedPoolPairData(subgraphPool *etypes.SubgraphPoolBase, tokenIn string, tokenOut string) *WeightedPoolPairData {
	wp := WeightedPoolPairData{}

	wp.SwapFee, _ = strconv.ParseFloat(subgraphPool.SwapFee, 64)

	for _, token := range subgraphPool.Tokens {
		b, _ := strconv.ParseFloat(token.Balance, 64)
		d := int(token.Decimals)
		w, _ := strconv.ParseFloat(token.Weight, 64)
		add := token.Address

		if strings.EqualFold(tokenIn, add) {
			wp.BalanceIn = b
			wp.DecimalsIn = d
			wp.WeightIn = w
		}
		if strings.EqualFold(tokenOut, add) {
			wp.BalanceOut = b
			wp.DecimalsOut = d
			wp.WeightOut = w
		}
	}
	return &wp
}
