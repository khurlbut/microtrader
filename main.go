package main

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/khurlbut/microtrader/bank"
	"github.com/khurlbut/microtrader/websocketbinance"
)

const (
	BTCUSDT = "btcusdt"
)

func main() {
	url := websocketbinance.Url([]string{BTCUSDT})
	account := bank.NewAccount(1000.00, 100.00)

	websocket, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer websocket.Close()

	for {
		_, message, err := websocket.ReadMessage()
		if err != nil {
			log.Fatal("read:", err)
		}

		price := websocketbinance.ToPrice(message)
		if price.IsActionable(account) {
			// Calculate purchase price (bid + (ask - bid) / 2) * random(0.95, 1.05)
			// Withdraw funds from account
			// Create purchase order
			// Place order on exchange API
		}

		// process executed buy orders
		// cancel expireed buy orders

		// process executed sell orders
	}
}
