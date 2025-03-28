package main

import (
	"fmt"
)

// ctrl + D
// ctrl + shift + l
// the interface{} type is a special type that can represent any other data type.
// It acts as a container for values of any type, providing flexibility but sacrificing some type safety and performance optimizations.
//*Node is a pointer to another Node. It allows the current Node to reference another Node in memory.

type Node struct {
	Value any
	Next  *Node
}

type LinkedList struct {
	Head *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{Head: nil, size: 0}
}

func (ll *LinkedList) PushFront(value any) {
	newNode := &Node{Value: value, Next: ll.Head}
	ll.Head = newNode
	ll.size++
}

func (ll *LinkedList) Append(value any) {
	newNode := &Node{Value: value, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
	} else {
		lastNode := ll.Head

		for lastNode.Next != nil {
			lastNode = lastNode.Next
		}

		lastNode.Next = newNode
	}

	ll.size++
}

func (ll *LinkedList) Print() {
	current := ll.Head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}
func (ll *LinkedList) SliceMy() []any {
	current := ll.Head
	var result []any

	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

func (ll *LinkedList) Length() int {
	return ll.size
}

func main() {
	list := NewLinkedList()
	list.PushFront("Saddle")
	list.Append("Whip")
	list.PushFront("Belt")
	list.Append("Pistol")
	list.Append(1900)
	list.Append(true)
	fmt.Println("List: ")
	list.Print()
	linkedListLength := list.Length()
	MySlice := list.SliceMy()
	fmt.Println(MySlice)
	fmt.Println("Length: ", linkedListLength)

}
