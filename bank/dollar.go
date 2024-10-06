package bank

import f "github.com/khurlbut/microtrader/float"

// ToDollarString converts a float64 to a string.
func ToDollarString(c float64) string {
	return f.ToStringWithPrecision(c, 2)
}
