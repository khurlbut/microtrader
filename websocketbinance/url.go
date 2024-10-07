package websocketbinance

import "strings"

const (
	WebSocketURL = "wss://stream.binance.us:9443"
)

func Url(symbols []string) string {
	return WebSocketURL + "/stream?streams=" + streams(symbols)
}

func streams(symbols []string) string {
	var streams []string
	for _, symbol := range symbols {
		streams = append(streams, symbol+"@ticker")
	}
	return strings.Join(streams, "/")
}
