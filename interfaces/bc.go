package main

import "errors"

type Bitcoin struct {
	name    string
	balance int
	fee     int
}

// NewBitcoinAccount si constructor for type Bitcoin
func NewBitcoinAccount() *Bitcoin {
	return &Bitcoin{
		name:    "bitcoin",
		balance: 0,
		fee:     300,
	}
}

func (b *Bitcoin) GetBalance() int {
	return b.balance
}

func (b *Bitcoin) Deposit(amount int) {
	b.balance += amount
}

func (b *Bitcoin) Withdraw(amount int) error {
	newBalance := b.balance - amount - b.fee
	if newBalance < 0 {
		return errors.New(insufficientFundsError)
	}
	b.balance = newBalance
	return nil
}

func (b *Bitcoin) GetName() string {
	return b.name
}
