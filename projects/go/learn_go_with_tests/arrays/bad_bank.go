package arrays

// Transaction represents a financial transaction between two parties, with a specified amount of money being transferred from one party to another.
type Transaction struct {
	From string
	To   string
	Sum  float64
}

// NewTransaction creates a new Transaction struct given the sender, receiver, and amount of money being transferred.
func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

// Account represents a bank account with a name and a balance.
type Account struct {
	Name    string
	Balance float64
}

// NewBalanceFor calculates the new balance for a given account based on a list of transactions.
func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(transactions, applyTransaction, account)
}

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}

// BalanceFor calculates the balance for a given name based on a list of transactions.
func BalanceFor(transactions []Transaction, name string) float64 {
	adjustBalance := func(currentBalance float64, t Transaction) float64 {
		if t.From == name {
			return currentBalance - t.Sum
		}
		if t.To == name {
			return currentBalance + t.Sum
		}
		return currentBalance
	}
	return Reduce(transactions, adjustBalance, 0.0)
}
