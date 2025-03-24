package main

import (
	"fmt"
)

func main() {
	type Doctor struct {
		FirstName string
		LastName  string
		Education string
		Married   bool
		Age       int
	}

	person := Doctor{
		FirstName: "Jim",
		LastName:  "Bukkakki",
		Education: "9 PHDs",
		Married:   false,
		Age:       89,
	}
	fmt.Println(person.FirstName, person.LastName, person.Education, person.Married, person.Age)

	// Pointers to structs are common for methods
	var personPtr *Doctor = &person
	fmt.Println(personPtr.LastName)

}
