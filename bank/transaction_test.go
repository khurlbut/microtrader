package bank

import (
	"testing"
)

// TestNewWithdrawlTransaction tests the newWithdrawlTransaction function.
func TestNewWithdrawlTransaction(t *testing.T) {
	transaction := newWithdrawlTransaction("12345", 100.00)

	if transaction.accountNumber != "12345" {
		t.Errorf("Expected 12345, got %v", transaction.accountNumber)
	}

	if transaction.Amount() != 100.00 {
		t.Errorf("Expected 100.00, got %v", transaction.Amount())
	}

	if transaction.Fulfilled() {
		t.Errorf("Expected false, got %v", transaction.Fulfilled())
	}
}

// TestTransactionNewDepositTransaction tests the Transaction.NewDepositTransaction method.
func TestTransactionNewDepositTransaction(t *testing.T) {
	transaction := newWithdrawlTransaction("54321", 100.00)
	deposit, err := transaction.NewDepositTransaction(100.00)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if deposit.accountNumber != "54321" {
		t.Errorf("Expected 54321, got %v", deposit.accountNumber)
	}

	if deposit.Amount() != 100.00 {
		t.Errorf("Expected 100.00, got %v", deposit.Amount())
	}

	if !transaction.Fulfilled() {
		t.Errorf("Expected true, got %v", transaction.Fulfilled())
	}

	_, err = transaction.NewDepositTransaction(99.99)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	_, err = transaction.NewDepositTransaction(200.00)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
