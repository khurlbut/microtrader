package bank

import "testing"

func TestNewTransactions(t *testing.T) {
	transactions := newTransactions("12345")
	totalOutstanding := ToDollarString(transactions.totalOutstanding())
	if totalOutstanding != "0.00" {
		t.Errorf("Expected 0.00, got %s", totalOutstanding)
	}
}

func TestAddTranscaction(t *testing.T) {
	transactions := newTransactions("12345")
	transaction := newWithdrawlTransaction("12345", 100.00)
	transactions.addTransaction(transaction)
	totalOutstanding := ToDollarString(transactions.totalOutstanding())
	if totalOutstanding != "100.00" {
		t.Errorf("Expected 100.00, got %s", totalOutstanding)
	}
}

func TestTransactionsOutstandingDoesNotIncludeFulfilledTransactions(t *testing.T) {
	transactions := newTransactions("12345")
	transaction := newWithdrawlTransaction("12345", 100.00)
	transaction.fulfill()
	transactions.addTransaction(transaction)
	totalOutstanding := ToDollarString(transactions.totalOutstanding())
	if totalOutstanding != "0.00" {
		t.Errorf("Expected 0.00, got %s", totalOutstanding)
	}
}

func TestAddTranscactionIncomingTransactionMustBelongToAccount(t *testing.T) {
	transactions := newTransactions("12345")
	transaction := newWithdrawlTransaction("12345", 100.00)
	err := transactions.addTransaction(transaction)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	transaction = newWithdrawlTransaction("54321", 100.00)
	err = transactions.addTransaction(transaction)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestAddTranscactionIncomingTransactionMustNotExist(t *testing.T) {
	transactions := newTransactions("12345")
	transaction := newWithdrawlTransaction("12345", 100.00)
	err := transactions.addTransaction(transaction)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	err = transactions.addTransaction(transaction)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestLookupTransaction(t *testing.T) {
	transactions := newTransactions("54321")

	transaction1 := newWithdrawlTransaction("54321", 100.00)
	transactions.addTransaction(transaction1)
	transaction2 := newWithdrawlTransaction("54321", 100.00)
	transactions.addTransaction(transaction2)
	transaction3 := newWithdrawlTransaction("54321", 100.00)
	transactions.addTransaction(transaction3)

	if transactions.lookupTransaction(transaction1.transactionId) != transaction1 {
		t.Errorf("Expected transaction1, got %v", transactions.lookupTransaction(transaction1.transactionId))
	}

	if transactions.lookupTransaction(transaction2.transactionId) != transaction2 {
		t.Errorf("Expected transaction2, got %v", transactions.lookupTransaction(transaction2.transactionId))
	}

	if transactions.lookupTransaction(transaction3.transactionId) != transaction3 {
		t.Errorf("Expected transaction3, got %v", transactions.lookupTransaction(transaction1.transactionId))
	}
}

func TestRemoveTransaction(t *testing.T) {
	transactions := newTransactions("54321")

	transaction1 := newWithdrawlTransaction("54321", 100.00)
	transactions.addTransaction(transaction1)
	transaction2 := newWithdrawlTransaction("54321", 100.00)
	transactions.addTransaction(transaction2)
	transaction3 := newWithdrawlTransaction("54321", 100.00)
	transactions.addTransaction(transaction3)

	// remove transaction2
	transactions.removeTransaction(transaction2.transactionId)

	if transactions.lookupTransaction(transaction1.transactionId) != transaction1 {
		t.Errorf("Expected transaction1, got %v", transactions.lookupTransaction(transaction1.transactionId))
	}

	// transaction2 should be removed
	if transactions.lookupTransaction(transaction2.transactionId) != nil {
		t.Errorf("Expected nil, got %v", transactions.lookupTransaction(transaction2.transactionId))
	}

	if transactions.lookupTransaction(transaction3.transactionId) != transaction3 {
		t.Errorf("Expected transaction3, got %v", transactions.lookupTransaction(transaction1.transactionId))
	}
}

func TestTotalAmount(t *testing.T) {
	transactions := newTransactions("12345")
	transaction := newWithdrawlTransaction("12345", 100.00)
	transactions.addTransaction(transaction)
	transaction = newWithdrawlTransaction("12345", 100.00)
	transactions.addTransaction(transaction)
	totalAmount := ToDollarString(transactions.totalOutstanding())
	if totalAmount != "200.00" {
		t.Errorf("Expected 200.00, got %s", totalAmount)
	}
}
