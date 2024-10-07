package pricetracker

import "testing"

func TestNewPrice(t *testing.T) {
	t.Skip("Test not implemented")
	t.Errorf("Test not implemented")
}

func TestBuyRecommended(t *testing.T) {
	p := Price{}
	if p.buyRecommended() != true {
		t.Errorf("Expected true, got false")
	}
}
