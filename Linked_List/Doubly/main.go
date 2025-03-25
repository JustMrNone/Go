package main

import (
	"fmt"
)

// Node represents a single node in the doubly linked list.
// It contains a value, a pointer to the next node, and a pointer to the previous node.
type Node struct {
	Value any   // The value stored in the node (can be of any type).
	Next  *Node // Pointer to the next node in the list.
	Prev  *Node // Pointer to the previous node in the list.
}

// LinkedList represents the doubly linked list.
// It contains pointers to the head and tail nodes, as well as the size of the list.
type LinkedList struct {
	Head *Node // Pointer to the first node in the list.
	Tail *Node // Pointer to the last node in the list.
	Size int   // The number of nodes in the list.
}

// NewList creates and returns a new, empty doubly linked list.
func NewList() *LinkedList {
	return &LinkedList{
		Head: nil, // Head starts as nil since the list is empty.
		Tail: nil, // Tail starts as nil since the list is empty.
		Size: 0,   // Size is initialized to 0.
	}
}

// PushFront adds a new node with the given value to the front of the list.
func (ll *LinkedList) PushFront(value any) {
	// Create a new node with the given value.
	// Its Next pointer will point to the current head, and Prev will be nil.
	newNode := &Node{Value: value, Next: ll.Head, Prev: nil}

	// If the list is not empty, update the current head's Prev pointer to the new node.
	if ll.Head != nil {
		ll.Head.Prev = newNode
	} else {
		// If the list is empty, the new node becomes the tail as well.
		ll.Tail = newNode
	}
	// Update the head pointer to the new node.
	ll.Head = newNode

	// Increment the size of the list.
	ll.Size++
}

// Append adds a new node with the given value to the end of the list.
func (ll *LinkedList) Append(value any) {
	// Create a new node with the given value.
	// Its Prev pointer will point to the current tail, and Next will be nil.
	newNode := &Node{Value: value, Prev: ll.Tail, Next: nil}

	// If the list is not empty, update the current tail's Next pointer to the new node.
	if ll.Tail != nil {
		ll.Tail.Next = newNode
	} else {
		// If the list is empty, the new node becomes the head as well.
		ll.Head = newNode
	}
	// Update the tail pointer to the new node.
	ll.Tail = newNode

	// Increment the size of the list.
	ll.Size++
}

// Print traverses the list from the head to the tail and prints the values of all nodes.
func (ll *LinkedList) Print() {
	current := ll.Head // Start from the head of the list.
	for current != nil {
		// Print the value of the current node.
		fmt.Print(current.Value)
		// If there is a next node, print the separator "<->".
		if current.Next != nil {
			fmt.Print(" <-> ")
		}
		// Move to the next node.
		current = current.Next
	}
	// Print a newline at the end.
	fmt.Println()
}

func (ll *LinkedList) IntoSlice() []any {
	current := ll.Head
	var sol []any
	for current != nil {
		sol = append(sol, current.Value)
		current = current.Next
	}
	return sol
}

// Length returns the number of nodes in the list.
func (ll *LinkedList) Length() int {
	return ll.Size
}

func main() {
	fmt.Println("Doubly Linked List:")

	// Create a new doubly linked list.
	List := NewList()

	// Add elements to the list.
	List.PushFront(123) // Add 123 to the front of the list.
	List.Append("None") // Add "None" to the end of the list.
	List.PushFront(5.5) // Add 5.5 to the front of the list.
	List.Append(true)   // Add true to the end of the list.

	// Print the contents of the list.
	List.Print()
	MySlice := List.IntoSlice()
	fmt.Println(MySlice...)
	// Print the length of the list.
	fmt.Printf("List Length: %v", List.Length())
}
