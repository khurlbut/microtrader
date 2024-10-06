package websocketbinance

import "strings"

func streamsParam(symbols []string) string {
	var streams []string
	for _, symbol := range symbols {
		streams = append(streams, symbol+"@ticker")
	}
	return strings.Join(streams, "/")
}
