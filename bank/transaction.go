package bank

import (
	"fmt"

	f "github.com/khurlbut/triangulator/float"
)

// Transaction represents a pseudo bank transaction.
type Transaction struct {
	transactionId string
	accountNumber string
	amount        float64
	fulfilled     bool
}

// Amount returns the amount of the transaction.
func (t *Transaction) Amount() float64 {
	return t.amount
}

// Fulfilled returns true if the transaction has been fulfilled.
func (t *Transaction) Fulfilled() bool {
	return t.fulfilled
}

func (t *Transaction) fulfill() {
	t.fulfilled = true
}

// newWithdrawlTransaction creates a new withdrawl transaction.
func newWithdrawlTransaction(accountNumber string, amount float64) *Transaction {
	transactionId, err := generateRandomID(16)
	if err != nil {
		panic(fmt.Sprintf("failed to generate random transaction ID: %v", err))
	}

	return &Transaction{
		transactionId: transactionId,
		accountNumber: accountNumber,
		amount:        amount,
	}
}

// NewDepositTransaction creates a new deposit transaction.
func (t *Transaction) NewDepositTransaction(amount float64) (*Transaction, error) {
	if t.fulfilled {
		return nil, fmt.Errorf("withdrawl transaction already fulfilled")
	}
	if amount < t.amount {
		return nil, fmt.Errorf("deposit amount is less than the transaction amount")
	}
	if amount >= t.amount*2 {
		return nil, fmt.Errorf("deposit amount is greater than twice the transaction amount")
	}
	t.fulfilled = true
	return &Transaction{
		accountNumber: t.accountNumber,
		amount:        amount,
	}, nil
}

// ToString converts a Transaction to a string.
func (t *Transaction) ToString() string {
	return f.ToStringWithPrecision(t.amount, 2)
}
