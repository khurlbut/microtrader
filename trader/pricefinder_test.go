package trader

import (
	"math"
	"testing"
)

// TestPriceFinder tests the Price method of the PriceFinder struct
func TestPriceFinder(t *testing.T) {
	testcases := []struct {
		name          string
		makerfee      float64
		takerfee      float64
		profitTarget  float64
		tradeAmount   float64
		currentPrice  float64
		expectedPrice float64
	}{
		{
			name:          "nothing",
			expectedPrice: 0.0,
		},
		{
			name:          "trade equals price at 1 dollar",
			tradeAmount:   1.0,
			currentPrice:  1.0,
			profitTarget:  0.10,
			expectedPrice: 1.10,
		},
		{
			name:          "trade equals price at 10 dollars",
			tradeAmount:   10.0,
			currentPrice:  10.0,
			profitTarget:  0.10,
			expectedPrice: 10.10,
		},
		{
			name:          "trade is 10 dollars, price is 100 dollars",
			tradeAmount:   10.0,
			currentPrice:  100.0,
			profitTarget:  0.10,
			expectedPrice: 101.00,
		},
		{
			name:          "trade is 100 dollars, price is 60000 dollars",
			tradeAmount:   100.0,
			currentPrice:  60000.0,
			profitTarget:  0.10,
			expectedPrice: 60060.00,
		},
		{
			name:          "trade equals price at 10 dollars with fees",
			tradeAmount:   10.0,
			currentPrice:  10.0,
			profitTarget:  0.10,
			makerfee:      0.001,
			takerfee:      0.001,
			expectedPrice: 10.12,
		},
		{
			name:          "trade is 100 dollars, price is 60000 dollars",
			tradeAmount:   100.0,
			currentPrice:  60000.0,
			profitTarget:  0.10,
			makerfee:      0.001,
			takerfee:      0.001,
			expectedPrice: 60180.00,
		},
		{
			name:          "trade is 100 dollars, price is 65000 dollars",
			tradeAmount:   100.0,
			currentPrice:  65000.0,
			profitTarget:  0.10,
			makerfee:      0.001,
			takerfee:      0.001,
			expectedPrice: 65195.00,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			pf := NewPriceFinder(tc.makerfee, tc.takerfee, tc.profitTarget)
			price := pf.Price(tc.tradeAmount, tc.currentPrice)
			if !almostEqual(price, tc.expectedPrice, 1e-9) {
				t.Errorf("expected %f, got %f", tc.expectedPrice, price)
			}
		})
	}
}

// almostEqual checks if two floats are equal within a tolerance
func almostEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}

func TestPriceFinderShowIncrease(t *testing.T) {
	pf := NewPriceFinder(0.001, 0.001, 0.10)
	tradeAmounts := []float64{50.0, 75.0, 100.0, 150.0, 200.0, 250.0, 500.0, 575.0, 1000.0, 5000.0, 10000.0}
	currentPrice := 65000.0
	for _, tradeAmount := range tradeAmounts {

		price := pf.Price(tradeAmount, currentPrice)
		increase := price - currentPrice
		t.Logf("Trade amount: %f, Price: %f, Increase: %f", tradeAmount, price, increase)
	}
}
