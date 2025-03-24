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

// Value vs Pointer: In Go, when you pass a struct (or any value type) by value (Page), a copy of that struct is made and passed to the function.
// Any changes made within the function do not affect the original struct.
// On the other hand, passing a pointer (*Page) allows the function to modify the original struct directly.

func (book Book) print_details() {
	fmt.Printf("Book %s was written by %s.", book.Name, book.Author)
	fmt.Printf("\nIt contains %d pages.\n", book.Pages)
}

//Using a pointer allows you to modify the original data directly, rather than working with a copy.
//This is particularly useful when you want to update the fields of a struct or avoid copying large amounts of data.
//the Updatedt method is defined with a receiver of type *Book

func (book *Book) Updatedt(name string, page int) {
	book.Name = name
	book.Pages = page
}

func main() {
	book1 := Book{"Crime and Punishment", "Fyodor Dostoevsky", 879}
	book1.print_details()
	book1.Updatedt("Brothers Karmazov", 1080)
	book1.print_details()
}
