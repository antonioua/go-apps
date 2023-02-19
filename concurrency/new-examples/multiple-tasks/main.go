package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	done := make(chan struct{})

	go task1(done)
	go task2(done)
	go task3(done)
	go task4(done)

	<-done
	<-done
	<-done
	<-done

	fmt.Println(time.Since(now))
	fmt.Println("main is done.")
}

func task1(done chan struct{}) {
	time.Sleep(time.Millisecond * 200)
	fmt.Println("task1")
	done <- struct{}{}
}

func task2(done chan struct{}) {
	time.Sleep(time.Millisecond * 200)
	fmt.Println("task2")
	done <- struct{}{}
}

func task3(done chan struct{}) {
	time.Sleep(time.Millisecond * 200)
	fmt.Println("task3")
	done <- struct{}{}
}

func task4(done chan struct{}) {
	time.Sleep(time.Millisecond * 200)
	fmt.Println("task4")
	done <- struct{}{}
}
