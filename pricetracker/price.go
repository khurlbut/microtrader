package pricetracker

import "github.com/khurlbut/microtrader/bank"

type Price struct {
	Price string
}

func (p *Price) IsActionable(a *bank.Account) bool {
	if a.CashAvailable() {
		return p.buyRecommended()
	}
	return false
}

func (p *Price) buyRecommended() bool {
	// evaluate things like variation and order frequency
	return true
}
