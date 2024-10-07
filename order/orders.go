package order

import "fmt"

type Orders struct {
	accountNumber string
	side          OrderSide
	orders        []Order
}

func NewOrders(side OrderSide, accountNumber string) *Orders {
	if !side.IsValid() {
		fmt.Printf("invalid order side %q. side must be 'buy' or 'sell'", side)
		return nil
	}

	return &Orders{
		accountNumber: accountNumber,
		side:          side.normalizeOrderSide(),
		orders:        make([]Order, 0),
	}
}

func (o *Orders) AddOrder(order Order) {
	if order.side != o.side {
		fmt.Printf("cannot add order to orders: order side %q does not match Orders side %q", order.side, o.side)
		return
	}

	o.orders = append(o.orders, order)
}
