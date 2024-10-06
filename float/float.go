package float

import (
	"math"
	"strconv"
)

const epsilon = 0.00000001

// Equal checks if two floating-point numbers are approximately equal.
func Equal(a, b float64) bool {
	return math.Abs(a-b) <= epsilon
}

// NotEqual checks if two floating-point numbers are not approximately equal.
func NotEqual(a, b float64) bool {
	return !Equal(a, b)
}

// ToString converts a float64 to a string.
func ToString(f float64) string {
	return ToStringWithPrecision(f, -1)
}

// ToStringWithPrecision converts a float64 to a string with specified precision.
func ToStringWithPrecision(f float64, precision int) string {
	return strconv.FormatFloat(f, 'f', precision, 64)
}
