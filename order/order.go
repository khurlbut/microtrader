package order

import (
	"fmt"

	"github.com/khurlbut/microtrader/bank"
	"github.com/khurlbut/microtrader/identity"
	"github.com/khurlbut/microtrader/pricetracker"
)

// OrderState represents the state of a PurchaseOrder.
type OrderState string

// Constants defining the allowed states for a PurchaseOrder.
const (
	StateNew      OrderState = "new"
	StatePlaced   OrderState = "placed"
	StateExecuted OrderState = "executed"
)

// Order represents a purchase order.
type Order struct {
	orderNumber string
	withdrawl   bank.Transaction
	// Exchange API order
	price  pricetracker.Price
	state  OrderState
	expire int64
}

// NewOrder creates a new order in the "new" state.
func NewOrder(withdrawl bank.Transaction, price pricetracker.Price) *Order {
	orderNumber, err := identity.GenerateRandomID(16)
	if err != nil {
		panic(fmt.Sprintf("failed to generate random purchase order number: %v", err))
	}

	return &Order{
		orderNumber: orderNumber,
		withdrawl:   withdrawl,
		price:       price,
		state:       StateNew,
	}
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
