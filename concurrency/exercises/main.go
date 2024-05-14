package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	go task1()
	go task2()
	go task3()
	go task4()
	fmt.Println("Time elapsed: ", time.Since(now))
	time.Sleep(time.Second * 2)
}

func task1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task1 ")
}

func task2() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("task2 ")
}

func task3() {
	time.Sleep(1 * time.Second)
	fmt.Println("task3 ")
}

func task4() {
	time.Sleep(2 * time.Second)
	fmt.Println("task4 ")
}
