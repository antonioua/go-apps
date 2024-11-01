package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	now := time.Now()
	defer func() {
		elapsed := time.Since(now)
		fmt.Printf("Time elapsed: %.2f seconds\n", elapsed.Seconds())
	}()

	ch := make(chan string)
	go throwNinjaStar(ch)

	for {
		msg, open := <-ch
		if !open {
			break
		}
		fmt.Println(msg)
	}

}

func throwNinjaStar(ch chan string) {
	numOfRounds := 3
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i <= numOfRounds; i++ {
		ch <- fmt.Sprintf("Ninja score is: %d", rnd.Intn(10))
		time.Sleep(1 * time.Second)
	}
	close(ch)
}
