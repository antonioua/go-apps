package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var balance sync.Mutex

type Income struct {
	Source string
	Amount float64
}

func main() {
	// variable for bank balance
	var bankBalance float64

	// print out starting values
	fmt.Printf("Starting balance is $%.2f\n", bankBalance)

	// define weekly revenue
	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Part time job", Amount: 100},
		{Source: "Prizes", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	wg.Add(len(incomes))

	// loop through 52 weeks and print out how much is made; keep a running total
	for _, income := range incomes {
		go func(income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				tmpBalance := bankBalance
				tmpBalance += income.Amount
				bankBalance = tmpBalance
				balance.Unlock()

				fmt.Printf("Source: %s, week %d, amount: $%.2f\n", income.Source, week, tmpBalance)
			}
		}(income)
	}

	wg.Wait()

	// print out final balance
	fmt.Printf("Total revenue: $%.2f\n", bankBalance)

}
