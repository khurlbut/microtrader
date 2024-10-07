package order

import (
	"fmt"

	"github.com/khurlbut/microtrader/bank"
	"github.com/khurlbut/microtrader/identity"
	"github.com/khurlbut/microtrader/pricetracker"
)

// Order represents a purchase order.
type Order struct {
	orderNumber string
	side        OrderSide
	withdrawl   bank.Transaction
	// Exchange API order
	price pricetracker.Price
	state OrderState
}

// NewOrder creates a new order in the "new" state.
func NewOrder(side OrderSide, withdrawl bank.Transaction, price pricetracker.Price) (*Order, error) {
	orderNumber, err := identity.GenerateRandomID(16)
	if err != nil {
		panic(fmt.Sprintf("failed to generate random purchase order number: %v", err))
	}

	if !side.IsValid() {
		return nil, fmt.Errorf("invalid order side %q. side must be 'buy' or 'sell'", side)
	}

	return &Order{
		orderNumber: orderNumber,
		side:        side,
		withdrawl:   withdrawl,
		price:       price,
		state:       StateNew,
	}, nil
}

// SetState changes the state of the purchase order, allowing only valid transitions.
func (po *Order) SetState(newState OrderState) error {
	switch newState {
	case StateNew, StatePlaced, StateExecuted:
		po.state = newState
		return nil
	default:
		return fmt.Errorf("invalid state transition to %s", newState)
	}
}
