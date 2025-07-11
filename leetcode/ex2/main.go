package main

// import all methods from structs package to be called without specifying pacakge name
import (
	. "github.com/antonioua/leetcode/ex2/structs"
	"strconv"
)

func main() {
	l1 := &LinkedList{}
	//l1.Prepend(&Node{Data: 3})
	//l1.Prepend(&Node{Data: 4})
	//l1.Prepend(&Node{Data: 2})
	l1.Prepend(&Node{Data: 9})
	l1.Prepend(&Node{Data: 9})
	l1.Prepend(&Node{Data: 9})
	l1.Prepend(&Node{Data: 9})
	l1.Prepend(&Node{Data: 9})
	l1.Prepend(&Node{Data: 9})
	l1.Prepend(&Node{Data: 9})

	l2 := &LinkedList{}
	//l2.Prepend(&Node{Data: 4})
	//l2.Prepend(&Node{Data: 6})
	//l2.Prepend(&Node{Data: 5})
	l2.Prepend(&Node{Data: 9})
	l2.Prepend(&Node{Data: 9})
	l2.Prepend(&Node{Data: 9})
	l2.Prepend(&Node{Data: 9})

	l1.PrintData()
	l2.PrintData()

	l3 := addTwoNumbers(l1, l2)
	l3.PrintData()
}

func addTwoNumbers(l1 *LinkedList, l2 *LinkedList) *LinkedList {

	var tmpList1 []int
	currentL1Head := l1.Head

	// Traverse the linkedList and collect the data into a slice
	for i := 0; i < l1.Length; i++ {
		tmpList1 = append(tmpList1, currentL1Head.Data)
		currentL1Head = currentL1Head.Next
	}

	var s1Reverse string
	// Rewrite the slice in reverse order
	for i := l1.Length - 1; i >= 0; i-- {
		s1Reverse += strconv.Itoa(tmpList1[i])
	}

	int1, _ := strconv.Atoi(s1Reverse)

	var tmpList2 []int
	currentL2Head := l2.Head

	// Traverse the linkedList and collect the data into a slice
	for i := 0; i < l2.Length; i++ {
		tmpList2 = append(tmpList2, currentL2Head.Data)
		currentL2Head = currentL2Head.Next
	}

	var s2Reverse string
	// Rewrite the slice in reverse order
	for i := l2.Length - 1; i >= 0; i-- {
		s2Reverse += strconv.Itoa(tmpList2[i])
	}

	int2, _ := strconv.Atoi(s2Reverse)

	sum := int1 + int2
	sumStr := strconv.Itoa(sum)
	var sumList []int

	for _, char := range sumStr {
		digit, _ := strconv.Atoi(string(char))
		sumList = append(sumList, digit)
	}

	l3 := &LinkedList{}

	for i := 0; i < len(sumList); i++ {
		l3.Prepend(&Node{Data: sumList[i]})
	}

	return l3
}
