package main

import (
	"fmt"
)

type Node struct {
	Data any
	Next *Node
}

type LinkedList struct {
	Head *Node
	Size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Size: 0,
	}
}

func (ll *LinkedList) Prepend(value any) {
	newNode := &Node{Data: value, Next: ll.Head}
	ll.Head = newNode
	ll.Size++
}

func (ll *LinkedList) Append(value any) {
	newNode := &Node{Data: value, Next: nil}
	if ll.Head == nil {
		ll.Head = newNode
	} else {
		lastNode := ll.Head

		for lastNode.Next != nil {
			lastNode = lastNode.Next
		}

		lastNode.Next = newNode
	}
	ll.Size++
}

func (ll *LinkedList) Print() {
	current := ll.Head
	for current != nil {
		fmt.Print(current.Data)
		if current.Next != nil {
			fmt.Print("->")
		}
		current = current.Next
	}

}

func (ll *LinkedList) IntoSlice() []any {
	current := ll.Head
	var ans []any

	for current != nil {
		ans = append(ans, current.Data)
		current = current.Next
	}
	return ans
}

func (ll *LinkedList) Length() int {
	return ll.Size
}

func main() {
	fmt.Println("Singly: ")
	Sll := NewLinkedList()
	if Sll == nil {
		fmt.Println("Failed to initialize LinkedList")
		return
	}
	Sll.Append(1)
	Sll.Append(2)
	Sll.Append(3)
	Sll.Print()
}
