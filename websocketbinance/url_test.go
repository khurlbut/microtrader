package websocketbinance

import "testing"

func TestUrl(t *testing.T) {
	testcases := []struct {
		pairs    []string
		expected string
	}{
		{[]string{"btcusdt"}, "wss://stream.binance.us:9443/stream?streams=btcusdt@ticker"},
		{[]string{"btcusdt", "ethusdt"}, "wss://stream.binance.us:9443/stream?streams=btcusdt@ticker/ethusdt@ticker"},
		{[]string{"btcusdt", "ethusdt", "solusdt"}, "wss://stream.binance.us:9443/stream?streams=btcusdt@ticker/ethusdt@ticker/solusdt@ticker"},
	}

	for _, tc := range testcases {
		actual := Url(tc.pairs)
		if actual != tc.expected {
			t.Errorf("Expected %q but got %q", tc.expected, actual)
		}
	}
}

func TestStreams(t *testing.T) {
	testcases := []struct {
		pairs    []string
		expected string
	}{
		{[]string{"btcusdt"}, "btcusdt@ticker"},
		{[]string{"btcusdt", "ethusdt"}, "btcusdt@ticker/ethusdt@ticker"},
		{[]string{"btcusdt", "ethusdt", "solusdt"}, "btcusdt@ticker/ethusdt@ticker/solusdt@ticker"},
	}

	for _, tc := range testcases {
		actual := streams(tc.pairs)
		if actual != tc.expected {
			t.Errorf("Expected %q but got %q", tc.expected, actual)
		}
	}
}
