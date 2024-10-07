package order

import "fmt"

type Orders struct {
	accountNumber string
	side          OrderSide
	orders        []*Order
}

func NewOrders(side OrderSide, accountNumber string) *Orders {
	if !side.IsValid() {
		fmt.Printf("invalid order side %q. side must be 'buy' or 'sell'", side)
		return nil
	}

	return &Orders{
		accountNumber: accountNumber,
		side:          side.normalizeOrderSide(),
		orders:        make([]*Order, 0),
	}
}

func (o *Orders) AddOrder(order *Order) {
	orderSide := order.side.normalizeOrderSide()
	if orderSide != o.side {
		fmt.Printf("cannot add order to orders: order side %q does not match Orders side %q.\n", order.side, o.side)
		return
	}

	if order.state != StateNew {
		fmt.Printf("cannot add order to orders: order state %q is not 'new'.\n", order.state)
		return
	}

	o.orders = append(o.orders, order)
}
