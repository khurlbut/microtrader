package bank

import "fmt"

// Transaction represents a list of pseudo bank transactions.
type Transactions struct {
	accountNumber string
	transactions  map[string]*Transaction
}

// newTransactions creates a new Transactions instance.
func newTransactions(accountNumber string) *Transactions {
	return &Transactions{
		accountNumber: accountNumber,
		transactions:  make(map[string]*Transaction),
	}

}

// addTransaction adds a transaction to the list.
func (t *Transactions) addTransaction(transaction *Transaction) error {
	if transaction.accountNumber != t.accountNumber {
		return fmt.Errorf("transaction does not belong to account")
	}
	if _, exists := t.transactions[transaction.transactionId]; exists {
		return fmt.Errorf("transaction already exists")
	}
	t.transactions[transaction.transactionId] = transaction
	return nil
}

// lookupTransaction looks up a transaction by ID.
func (t *Transactions) lookupTransaction(transactionId string) *Transaction {
	return t.transactions[transactionId]
}

// removeTransaction removes a transaction from the list.
func (t *Transactions) removeTransaction(transactionId string) {
	delete(t.transactions, transactionId)
}

// totalOutstanding returns the total amount of all transactions.
func (t *Transactions) totalOutstanding() float64 {
	total := 0.0
	for _, transaction := range t.transactions {
		if !transaction.Fulfilled() {
			total += transaction.Amount()
		}
	}
	return total
}
