package order

import "strings"

// OrderState represents the state of an Order.
type OrderState string

// Constants defining the allowed states for an Order.
const (
	StateNew      OrderState = "new"
	StatePlaced   OrderState = "placed"
	StateExecuted OrderState = "executed"
)

// OrderSide represents the side of a purchase order.
type OrderSide string

// Constants defining the allowed sides for an Order.
const (
	OrderSideBuy  OrderSide = "buy"
	OrderSideSell OrderSide = "sell"
)

func (s OrderSide) IsValid() bool {
	ns := s.normalizeOrderSide()
	return ns == OrderSideBuy || ns == OrderSideSell
}

func (s OrderSide) normalizeOrderSide() OrderSide {
	return OrderSide(strings.ToLower(string(s)))
}

func (s OrderSide) String() string {
	return string(s)
}
