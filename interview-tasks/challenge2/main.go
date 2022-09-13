package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for v := range ch {
			fmt.Println("From routine:", v)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("Before:", i)
		ch <- i
		fmt.Println("After:", i)
	}

	close(ch)
}
