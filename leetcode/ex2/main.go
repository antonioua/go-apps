package main

// import all methods from structs package to be called without specifying pacakge name
import (
	. "github.com/antonioua/leetcode/ex2/structs"
)

func main() {
	myList := &LinkedList{}
	myList.Prepend(&Node{Data: 1})
	myList.Prepend(&Node{Data: 2})
	myList.Prepend(&Node{Data: 9})
	myList.Prepend(&Node{Data: 3})
	myList.Prepend(&Node{Data: 7})
	myList.PrintData()
	myList.DeleteWithValue(3)
	myList.PrintData()
}
