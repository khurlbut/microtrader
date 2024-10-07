package websocketbinance

import (
	"encoding/json"
	"log"

	"github.com/khurlbut/microtrader/pricetracker"
)

func ToPrice(message []byte) *pricetracker.Price {
	responseData := unmarshal(message)
	return responseData.price()
}

func (c *ResponseData) price() *pricetracker.Price {
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
