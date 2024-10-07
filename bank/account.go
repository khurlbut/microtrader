package bank

import (
	"fmt"
	"strconv"

	"github.com/khurlbut/microtrader/identity"
)

// Account represents a pseudo bank account.
type Account struct {
	accountNumber         string
	initialAmount         float64
	transactionAmount     float64
	currentAmount         float64
	withdrawlTransactions *Transactions
}

// NewAccount creates a new Cash instance.
func NewAccount(InitialAmount float64, transactionAmount float64) *Account {
	accountNumber, err := identity.GenerateRandomID(16)
	if err != nil {
		panic(fmt.Sprintf("failed to generate random account number: %v", err))
	}

	return &Account{
		accountNumber:         accountNumber,
		initialAmount:         InitialAmount,
		transactionAmount:     transactionAmount,
		currentAmount:         InitialAmount,
		withdrawlTransactions: newTransactions(accountNumber),
	}
}

// NewAccountFromStrings creates a new Account instance from string inputs.
func NewAccountFromStrings(InitialAmount string, transactionAmount string) (*Account, error) {
	InitialAmountFloat, err := strconv.ParseFloat(InitialAmount, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid value for total: %v", err)
	}

	transactionAmountFloat, err := strconv.ParseFloat(transactionAmount, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid value for transactionAmount: %v", err)
	}

	return NewAccount(InitialAmountFloat, transactionAmountFloat), nil
}

// InitialAmount returns the initial amount of the account.
func (a *Account) InitialAmount() float64 {
	return a.initialAmount
}

// TransactionAmount returns the transaction amount of the account.
func (a *Account) TransactionAmount() float64 {
	return a.transactionAmount
}

// CurrentAmount returns the current amount of the account.
func (a *Account) CurrentAmount() float64 {
	return a.currentAmount
}

// CashAvailable returns true if the account has enough cash to make a transaction.
func (a *Account) CashAvailable() bool {
	return a.currentAmount >= a.transactionAmount
}

// Withdraw removes the transaction amount from the account.
func (a *Account) Withdraw() (*Transaction, error) {
	if a.CashAvailable() {
		a.currentAmount -= a.transactionAmount
		transaction := newWithdrawlTransaction(a.accountNumber, a.transactionAmount)
		a.withdrawlTransactions.addTransaction(transaction)

		return transaction, nil
	}

	return nil, fmt.Errorf("insufficient funds")
}

// Deposit adds the transaction amount to the account.
func (a *Account) Deposit(t *Transaction) error {
	if t.accountNumber != a.accountNumber {
		return fmt.Errorf("invalid account number")
	}

	if t.fulfilled {
		return fmt.Errorf("deposit transaction already fulfilled")
	}

	a.currentAmount += t.Amount()
	t.fulfill()
	a.withdrawlTransactions.removeTransaction(t.transactionId)

	return nil
}

// CalculateProfit returns the profit of the account.
func (a *Account) CalculateProfit() float64 {
	return a.currentAmount + a.withdrawlTransactions.totalOutstanding() - a.initialAmount
}
