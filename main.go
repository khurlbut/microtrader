package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/khurlbut/microtrader/websocketbinance"
)

const (
	BTCUSDT = "btcusdt"
)

func main() {
	url := websocketbinance.Url([]string{BTCUSDT})

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

		priceupdate := websocketbinance.ToPriceUpdate(message)
		fmt.Println(priceupdate)
	}
}
