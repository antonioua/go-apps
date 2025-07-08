package structs

import "fmt"

type Node struct {
	Next *Node
	Data int
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (l *LinkedList) Prepend(n *Node) {
	// Save the pointer to the old Node
	oldHead := l.Head
	// Change the head of linkedList to point to the new Node
	l.Head = n
	// Point the new Node to the old Node
	n.Next = oldHead
	l.Length++
}

func (l LinkedList) PrintData() {
	if l.Length == 0 {
		return
	}

	currentNode := l.Head

	for i := 0; i < l.Length; i++ {
		if i == l.Length-1 {
			fmt.Printf("%d\n", currentNode.Data)
			return
		}
		fmt.Printf("%d -> ", currentNode.Data)
		currentNode = currentNode.Next
	}
}

func (l *LinkedList) DeleteWithValue(v int) {
	if l.Head == nil {
		return
	}

	// If the first node
	if l.Head.Data == v {
		l.Head = l.Head.Next
		l.Length--
		return
	}

	// Loop through all the remaining nodes
	currentNode := l.Head

	for currentNode.Next != nil {
		if currentNode.Next.Data == v {
			currentNode.Next = currentNode.Next.Next
			l.Length--
			return
		}
		currentNode = currentNode.Next
	}
}
