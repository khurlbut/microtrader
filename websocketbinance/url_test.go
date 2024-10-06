package websocketbinance

import "testing"

func TestStreamsParam(t *testing.T) {
	testcases := []struct {
		pairs    []string
		expected string
	}{
		{[]string{"btcusdt"}, "btcusdt@ticker"},
		{[]string{"btcusdt", "ethusdt"}, "btcusdt@ticker/ethusdt@ticker"},
		{[]string{"btcusdt", "ethusdt", "solusdt"}, "btcusdt@ticker/ethusdt@ticker/solusdt@ticker"},
	}

	for _, tc := range testcases {
		actual := streamsParam(tc.pairs)
		if actual != tc.expected {
			t.Errorf("Expected %q but got %q", tc.expected, actual)
		}
	}
}
