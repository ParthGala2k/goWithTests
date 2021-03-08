package main

import (
	"fmt"
	"errors"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("Insufficient funds, cannot withdraw. Aborting")
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.Balance() {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
