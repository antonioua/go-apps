package example1

import "errors"

type WellsFargo struct {
	name    string
	balance int
}

// NewWellsFargoAccount is a constructor for type WellsFargo
func NewWellsFargoAccount() *WellsFargo {
	return &WellsFargo{
		name:    "wellfargo",
		balance: 0,
	}
}

func (w *WellsFargo) GetBalance() int {
	return w.balance
}

func (w *WellsFargo) Deposit(amount int) {
	w.balance += amount
}

func (w *WellsFargo) Withdraw(amount int) error {
	newBalance := w.balance - amount
	if newBalance < 0 {
		return errors.New(insufficientFundsError)
	}
	w.balance = newBalance
	return nil
}

func (w *WellsFargo) GetName() string {
	return w.name
}
