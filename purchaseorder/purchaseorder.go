package purchaseorder

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

// PurchaseOrder represents a purchase order.
type PurchaseOrder struct {
	PurchaseOrderNumber string
	Withdrawl           bank.Transaction
	// Exchange API order
	Price pricetracker.Price
	State OrderState
}

// NewPurchaseOrder creates a new purchase order in the "new" state.
func NewPurchaseOrder(withdrawl bank.Transaction, price pricetracker.Price) *PurchaseOrder {
	purchaseOrderNumber, err := identity.GenerateRandomID(16)
	if err != nil {
		panic(fmt.Sprintf("failed to generate random purchase order number: %v", err))
	}

	return &PurchaseOrder{
		PurchaseOrderNumber: purchaseOrderNumber,
		Withdrawl:           withdrawl,
		Price:               price,
		State:               StateNew,
	}
}

// SetState changes the state of the purchase order, allowing only valid transitions.
func (po *PurchaseOrder) SetState(newState OrderState) error {
	switch newState {
	case StateNew, StatePlaced, StateExecuted:
		po.State = newState
		return nil
	default:
		return fmt.Errorf("invalid state transition to %s", newState)
	}
}
