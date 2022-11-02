package math

import "math"

type WeightedPoolPairData struct {
	BalanceIn   float64 `json:"balanceIn"`
	BalanceOut  float64 `json:"balanceOut"`
	WeightIn    float64 `json:"weightIn"`
	WeightOut   float64 `json:"weightOut"`
	SwapFee     float64 `json:"swapFee"`
	DecimalsIn  int     `json:"decimalsIn"`
	DecimalsOut int     `json:"decimalsOut"`
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
