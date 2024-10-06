package websocketbinance

import (
	"encoding/json"
	"log"

	"github.com/khurlbut/microtrader/pricetracker"
)

func ToPriceUpdate(message []byte) *pricetracker.PriceUpdate {
	responseData := unmarshal(message)
	return responseData.priceUpdate()
}

func (c *ResponseData) priceUpdate() *pricetracker.PriceUpdate {
	return nil
}

func unmarshal(message []byte) ResponseData {
	var responseData ResponseData
	err := json.Unmarshal([]byte(message), &responseData)
	if err != nil {
		log.Fatalf("Error unmarshaling message: %v", err)
	}
	return responseData
}
