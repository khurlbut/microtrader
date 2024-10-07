package order

import (
	"testing"

	"github.com/khurlbut/microtrader/bank"
	"github.com/khurlbut/microtrader/pricetracker"
)

func TestNewOrders(t *testing.T) {
	testcases := []struct {
		side          OrderSide
		accountNumber string
	}{
		{OrderSideBuy, "123"},
		{OrderSideSell, "123"},
	}
	for _, tc := range testcases {
		orders := NewOrders(tc.side, tc.accountNumber)
		if orders.accountNumber != "123" {
			t.Errorf("Expected 123, got %s", orders.accountNumber)
		}
		if orders.side != tc.side {
			t.Errorf("Expected %q, got %q", tc.side, orders.side)
		}
	}
}

func TestAddOrder(t *testing.T) {
	testcases := []struct {
		side OrderSide
	}{
		{OrderSideBuy},
		{OrderSideSell},
	}
	for _, tc := range testcases {
		orders := NewOrders(tc.side, "123")
		order, err := NewOrder(tc.side, bank.Transaction{}, pricetracker.Price{})
		if err != nil {
			t.Errorf("unexpected error creating new order: %v", err)
		}
		orders.AddOrder(order)
		if len(orders.orders) != 1 {
			t.Errorf("Expected 1, got %d", len(orders.orders))
		}
	}
}

func TestAddOrder_invalidOrderSide(t *testing.T) {
	testcases := []struct {
		sideOrders OrderSide
		sideOrder  OrderSide
	}{
		{OrderSideBuy, OrderSideSell},
		{OrderSideSell, OrderSideBuy},
	}
	for _, tc := range testcases {
		orders := NewOrders(tc.sideOrders, "123")
		order, err := NewOrder(tc.sideOrder, bank.Transaction{}, pricetracker.Price{})
		if err != nil {
			t.Errorf("unexpected error creating new order: %v", err)
		}
		orders.AddOrder(order)
		if len(orders.orders) != 0 {
			t.Errorf("Expected 0, got %d", len(orders.orders))
		}
	}
}

func TestAddOrder_orderMustBeNew(t *testing.T) {
	testcases := []struct {
		state          OrderState
		expectedOrders int
	}{
		{StateNew, 1},
		{StatePlaced, 0},
		{StateExecuted, 0},
	}

	for _, tc := range testcases {
		orders := NewOrders(OrderSideBuy, "123")
		order, err := NewOrder(OrderSideBuy, bank.Transaction{}, pricetracker.Price{})
		if err != nil {
			t.Errorf("unexpected error creating new order: %v", err)
		}
		order.SetState(tc.state)
		orders.AddOrder(order)
		if len(orders.orders) != tc.expectedOrders {
			t.Errorf("Expected %d, got %d", tc.expectedOrders, len(orders.orders))
		}
	}
}
