package order

import (
	"testing"
)

func TestIsValidOrderSide(t *testing.T) {
	tests := []struct {
		name string
		side OrderSide
		want bool
	}{
		{
			name: "Buy",
			side: OrderSideBuy,
			want: true,
		},
		{
			name: "Sell",
			side: OrderSideSell,
			want: true,
		},
		{
			name: "Invalid",
			side: OrderSide("invalid"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.side.IsValid(); got != tt.want {
				t.Errorf("OrderSide.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderSideEquality(t *testing.T) {
	testcases := []struct {
		side1 OrderSide
		side2 OrderSide
		want  bool
	}{
		{OrderSideBuy, OrderSideBuy, true},
		{OrderSideSell, OrderSideSell, true},
		{OrderSideBuy, OrderSideSell, false},
		{OrderSideSell, OrderSideBuy, false},
	}
	for _, tc := range testcases {
		if tc.side1 == tc.side2 != tc.want {
			t.Errorf("Expected %v, got %v", tc.want, tc.side1 == tc.side2)
		}
	}
}

func TestOrderSideString(t *testing.T) {
	tests := []struct {
		name string
		o    OrderSide
		want string
	}{
		{
			name: "Buy",
			o:    OrderSideBuy,
			want: "buy",
		},
		{
			name: "Sell",
			o:    OrderSideSell,
			want: "sell",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.String(); got != tt.want {
				t.Errorf("OrderSide.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
