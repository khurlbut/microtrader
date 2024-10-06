package bank

import (
	"testing"
)

// TestAccount tests the Account struct.
func TestAccount(t *testing.T) {
	testcases := []struct {
		initialAmount     string
		transactionAmount string
		cashAvailable     bool
	}{
		{"1000.00", "100.00", true},
		{"10000.00", "100.00", true},
		{"10000.00", "1000.00", true},
		{"10000.00", "10000.00", true},
		{"1000.00", "10000.00", false},
	}

	for _, tc := range testcases {
		a, err := NewAccountFromStrings(tc.initialAmount, tc.transactionAmount)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		initialAmount := ToDollarString(a.InitialAmount())
		if initialAmount != tc.initialAmount {
			t.Errorf("Expected InitialAmount %q, got %q", tc.initialAmount, initialAmount)
		}

		currentAmount := ToDollarString(a.CurrentAmount())
		if currentAmount != tc.initialAmount {
			t.Errorf("Expected CurrentAmount %q, got %q", tc.initialAmount, currentAmount)
		}

		transactionAmount := ToDollarString(a.TransactionAmount())
		if transactionAmount != tc.transactionAmount {
			t.Errorf("Expected TransactionAmount %q, got %q", tc.transactionAmount, transactionAmount)
		}

		if a.CashAvailable() != tc.cashAvailable {
			t.Errorf("Expected CashAvailable %t, got %t", tc.cashAvailable, a.CashAvailable())
		}
	}
}

// TestAccountWithdrawl tests the Account.Withdraw method.
func TestAccountWithdrawl(t *testing.T) {
	testcases := []struct {
		initialAmount         string
		transactionAmount     string
		expectedCurrentAmount string
		cashAvailable         bool
	}{
		{"1000.00", "100.00", "900.00", true},
		{"10000.00", "1000.00", "9000.00", true},
		{"10000.00", "10000.00", "0.00", false},
	}

	for _, tc := range testcases {
		a, _ := NewAccountFromStrings(tc.initialAmount, tc.transactionAmount)

		withdrawl, err := a.Withdraw()
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if withdrawl.ToString() != tc.transactionAmount {
			t.Errorf("Expected TransactionAmount %q, got %q", tc.transactionAmount, withdrawl.ToString())
		}

		currentAmount := ToDollarString(a.CurrentAmount())
		if currentAmount != tc.expectedCurrentAmount {
			t.Errorf("Expected CurrentAmount %q, got %q", tc.expectedCurrentAmount, currentAmount)
		}

		if a.CashAvailable() != tc.cashAvailable {
			t.Errorf("Expected CashAvailable %t, got %t", tc.cashAvailable, a.CashAvailable())
		}
	}
}

// TestAccountDeposit tests the Account.Deposit method.
func TestAccountDeposit(t *testing.T) {
	a, _ := NewAccountFromStrings("1000", "100")

	// a deposit can only be performed after a withdrawl
	withdrawl, err := a.Withdraw()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	// a deposit transaction can only be created from a withdrawl transaction
	deposit, err := withdrawl.NewDepositTransaction(100.10)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	a.Deposit(deposit)

	currentAmount := ToDollarString(a.CurrentAmount())
	if currentAmount != "1000.10" {
		t.Errorf("Expected 1000.10, got %q", currentAmount)
	}

	profit := ToDollarString(a.CalculateProfit())
	if profit != "0.10" {
		t.Errorf("Expected 0.10, got %q", profit)
	}
}

// TestAccountDeposit tests the Account.Deposit method.
func TestAccountBadDeposit(t *testing.T) {
	a, _ := NewAccountFromStrings("1000", "100")

	withdrawl, err := a.Withdraw()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	deposit, err := withdrawl.NewDepositTransaction(100.10)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	deposit.accountNumber = "bad account number"
	err = a.Deposit(deposit)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
