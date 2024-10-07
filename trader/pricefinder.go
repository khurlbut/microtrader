package trader

type priceFinder struct {
	makerFee     float64
	takerFee     float64
	profitTarget float64
}

func NewPriceFinder(makerFee, takerFee, profitTarget float64) *priceFinder {
	return &priceFinder{
		makerFee:     makerFee,
		takerFee:     takerFee,
		profitTarget: profitTarget,
	}
}

func (pf *priceFinder) Price(tradeAmount, currentPrice float64) float64 {
	if tradeAmount == 0 || currentPrice == 0 {
		return 0.0
	}

	currentValue := tradeAmount * currentPrice
	fees := tradeAmount * (pf.makerFee + pf.takerFee)
	targetValue := currentValue + ((pf.profitTarget + fees) * currentPrice)

	return targetValue / tradeAmount
}
