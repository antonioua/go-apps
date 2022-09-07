package main

import (
	"fmt"
	"time"
)

const insufficientFundsError = "insufficient funds"

type IBankAccount interface {
	GetBalance() int
	Deposit(amount int) // 100 = 1 dollar
	Withdraw(amount int) error
	GetName() string
}

func main() {
	now := time.Now()
	myAccounts := []IBankAccount{
		NewWellsFargoAccount(),
		NewBitcoinAccount(),
	}

	for _, account := range myAccounts {
		account.Deposit(300)
		if err := account.Withdraw(50); err != nil {
			fmt.Printf("Err for account %s: %v\n", account.GetName(), err)
		}

		fmt.Printf("Balance for account %s: %d\n", account.GetName(), account.GetBalance())
		fmt.Println("Elapsed: ", time.Since(now))
	}
}
