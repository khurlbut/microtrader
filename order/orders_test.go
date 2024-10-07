package order

import (
	"testing"
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
		order := Order{
			side: tc.side,
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
		order := Order{
			side: tc.sideOrder,
		}
		orders.AddOrder(order)
		if len(orders.orders) != 0 {
			t.Errorf("Expected 0, got %d", len(orders.orders))
		}
	}
}
