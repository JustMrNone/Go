package main

import (
	"fmt"
)

type Book struct {
	Name   string
	Author string
	Pages  int
}

// Object-Oriented Programming (OOP) in Go is different from traditional OOP languages like Java or C++. Go doesn’t have classes, inheritance,
// or traditional objects, but it supports key OOP principles through structs, methods, interfaces, and composition.
// (book Book): This part before the function name (print_details) is a method receiver in Go.
// book is the receiver variable, meaning that within this method, book represents an instance of the Book struct.
// Book is the type associated with the method. This means print_details() is a method that belongs to the Book struct.

func (book Book) print_details() {
	fmt.Printf("Book %s was written by %s.", book.Name, book.Author)
	fmt.Printf("\nIt contains %d pages.\n", book.Pages)
}

func main() {
	book1 := Book{"Crime and Punishment", "Fyodor Dostoevsky", 879}
	book1.print_details()
	book1.Name = "Brothers Karmazov"
	book1.Pages = 1045
	book1.print_details()
}
