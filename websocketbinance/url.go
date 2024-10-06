package websocketbinance

import "strings"

func url(symbols []string) string {
	return WebSocketURL + "/stream?streams=" + streams(symbols)
}

func streams(symbols []string) string {
	var streams []string
	for _, symbol := range symbols {
		streams = append(streams, symbol+"@ticker")
	}
	return strings.Join(streams, "/")
}
