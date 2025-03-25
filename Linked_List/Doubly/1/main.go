package main

import (
	"fmt"
)

type Node struct {
	Data any
	Next *Node
	Prev *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func NewList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

func (ll *LinkedList) Prepend(value any) {
	NewNode := &Node{Data: value, Next: ll.Head, Prev: nil}
	if ll.Head != nil {
		ll.Head.Prev = NewNode
	} else {
		ll.Tail = NewNode
	}
	ll.Head = NewNode
	ll.Size++

}

func (ll *LinkedList) Append(value any) {
	NewNode := &Node{Data: value, Next: nil, Prev: ll.Tail}
	if ll.Tail != nil {
		ll.Tail.Next = NewNode
	} else {
		ll.Head = NewNode
	}
	ll.Tail = NewNode
	ll.Size++

}

func (ll *LinkedList) Print() {
	current := ll.Head
	for current != nil {
		fmt.Print(current.Data)

		if current.Next != nil {
			fmt.Print(" <-> ")
		}

		current = current.Next
	}

	fmt.Println()

}

func (ll *LinkedList) IntoTheSlice() []any {
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
	fmt.Println("Doubly: ")
	Dll := NewList()
	Dll.Append(1)
	Dll.Prepend(0)
	Dll.Prepend(-1)
	Dll.Append(2)
	MySlice := Dll.IntoTheSlice()
	fmt.Println(MySlice[0])

	Dll.Print()
}
