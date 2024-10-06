package websocketbinance

// ResponseData represents the wrapper for the TickerData
type ResponseData struct {
	Stream string     `json:"stream"` // Stream name
	Data   TickerData `json:"data"`   // Data part of the message
}

// TickerData represents the structure for the JSON data provided by the Binance WebSocket API
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
