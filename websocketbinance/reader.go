package websocketbinance

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gorilla/websocket"
)

// TickerData represents the structure for the provided JSON data
type TickerData struct {
	EventType        string `json:"e"` // Event type
	EventTime        int64  `json:"E"` // Event time
	Symbol           string `json:"s"` // Symbol (e.g., BTCUSDT)
	PriceChange      string `json:"p"` // Price change
	PriceChangePct   string `json:"P"` // Price change percentage
	WeightedAvgPrice string `json:"w"` // Weighted average price
	PrevClosePrice   string `json:"x"` // Previous close price
	LastPrice        string `json:"c"` // Last price
	LastQty          string `json:"Q"` // Last quantity
	BidPrice         string `json:"b"` // Bid price
	BidQty           string `json:"B"` // Bid quantity
	AskPrice         string `json:"a"` // Ask price
	AskQty           string `json:"A"` // Ask quantity
	OpenPrice        string `json:"o"` // Open price
	HighPrice        string `json:"h"` // High price
	LowPrice         string `json:"l"` // Low price
	Volume           string `json:"v"` // Volume
	QuoteVolume      string `json:"q"` // Quote volume
	OpenTime         int64  `json:"O"` // Open time
	CloseTime        int64  `json:"C"` // Close time
	FirstTradeID     int64  `json:"F"` // First trade ID
	LastTradeID      int64  `json:"L"` // Last trade ID
	NumTrades        int64  `json:"n"` // Number of trades
}

// CombinedStreamMessage represents the wrapper for the TickerData
type CombinedStreamMessage struct {
	Stream string     `json:"stream"` // Stream name
	Data   TickerData `json:"data"`   // Data part of the message
}

// ConnectWebSocket connects to the Binance WebSocket API for the given symbols
func ConnectWebSocket(symbols []string) {
	url := "wss://stream.binance.us:9443/stream?streams=" + createCombinedStream(symbols)

	fmt.Println("Connecting to URL:", url)
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatal("read:", err)
		}

		unmarshal(message)
	}
}

// createCombinedStream creates a combined stream URL for the given symbols
func createCombinedStream(symbols []string) string {
	fmt.Println("Creating combined stream for symbols:", symbols)
	// Example: btcusdt@ticker/ethusdt@ticker/solbtc@ticker
	var streams []string
	for _, symbol := range symbols {
		streams = append(streams, symbol+"@ticker")
	}
	return strings.Join(streams, "/")
}

func unmarshal(wrappedData []byte) CombinedStreamMessage {
	var combinedStreamMessage CombinedStreamMessage
	err := json.Unmarshal([]byte(wrappedData), &combinedStreamMessage)
	if err != nil {
		log.Fatalf("Error unmarshaling message: %v", err)
	}
	return combinedStreamMessage
}
