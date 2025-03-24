package main

type Node struct {
	Data interface{}
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{Head: nil}
}

func main() {

}
