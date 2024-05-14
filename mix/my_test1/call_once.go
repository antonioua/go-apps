package main

import "fmt"

func func1() {
	fmt.Println("func1")
}

func func2() {
	fmt.Println("func2")
}

func func3() {
	fmt.Println("func3")
}

func main() {
	func1Called := false
	for {
		if !func1Called {
			func1()
			func1Called = true
		}
		func2()
		func3()
	}
}
