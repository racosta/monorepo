// Package pointers provides a simple implementation of a Bitcoin wallet with deposit, withdraw, and balance check functionalities.
package pointers

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds is an error returned when a withdraw operation cannot be completed due to insufficient funds.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Bitcoin is a custom type representing the amount of Bitcoin.
type Bitcoin int

// String returns a string representation of the Bitcoin amount.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet is a struct representing a Bitcoin wallet.
type Wallet struct {
	balance Bitcoin
}

// Deposit adds the specified amount to the wallet's balance.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns the current balance of the wallet.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw subtracts the specified amount from the wallet's balance if sufficient funds are available.
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
