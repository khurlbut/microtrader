package order

import (
	"strings"
	"testing"

	"github.com/khurlbut/microtrader/bank"
	"github.com/khurlbut/microtrader/pricetracker"
)

func TestNewOrder(t *testing.T) {
	po, err := NewOrder(OrderSideBuy, bank.Transaction{}, pricetracker.Price{})
	if err != nil {
		t.Errorf("unexpected error creating new order: %v", err)
	}

	if po.state != StateNew {
		t.Errorf("expected initial state to be %s, got %s", StateNew, po.state)
	}

	if po.orderNumber == "" {
		t.Error("expected purchase order number to be generated")
	}
}

func TestSetState_ValidTransitions(t *testing.T) {
	po, err := NewOrder(OrderSideSell, bank.Transaction{}, pricetracker.Price{})
	if err != nil {
		t.Errorf("unexpected error creating new order: %v", err)
	}

	validStates := []OrderState{StateNew, StatePlaced, StateExecuted}

	for _, state := range validStates {
		if err := po.SetState(state); err != nil {
			t.Errorf("unexpected error setting state to %s: %v", state, err)
		}

		if po.state != state {
			t.Errorf("expected state to be %s, got %s", state, po.state)
		}
	}
}

func TestSetState_InvalidTransition(t *testing.T) {
	po, err := NewOrder("BUY", bank.Transaction{}, pricetracker.Price{})
	if err != nil {
		t.Errorf("unexpected error creating new order: %v", err)
	}

	invalidState := OrderState("invalidState")

	if err := po.SetState(invalidState); err == nil {
		t.Errorf("expected an error when setting an invalid state, but got none")
	}
}

func TestSideMustBeOrderValid(t *testing.T) {
	for _, side := range []OrderSide{"BUY", "SELL"} {
		if _, err := NewOrder(side, bank.Transaction{}, pricetracker.Price{}); err != nil {
			t.Errorf("unexpected error creating new order: %v", err)
		}

		side = OrderSide(strings.ToLower(string(side)))
		if _, err := NewOrder(side, bank.Transaction{}, pricetracker.Price{}); err != nil {
			t.Errorf("unexpected error creating new order: %v", err)
		}
	}

	if _, err := NewOrder("INVALID", bank.Transaction{}, pricetracker.Price{}); err == nil {
		t.Errorf("expected an error when setting an invalid side, but got none")
	}
}
