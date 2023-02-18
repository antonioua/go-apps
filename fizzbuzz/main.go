package main

import "fmt"

func main() {
	var res []string
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			res = append(res, "FizzBuzz")
			// The continue statement is used when you want to skip the remaining portion of the loop,
			//and return to the top of the loop and continue a new iteration.
			continue
		}
		if i%3 == 0 {
			res = append(res, "Fizz")
			continue
		}
		if i%5 == 0 {
			res = append(res, "Buzz")
			continue
		}
		res = append(res, fmt.Sprintf("%v", i))
	}

	fmt.Println(res)
}
