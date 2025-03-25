package main

import (
	"fmt"
)

func main() {
	// when you initialze a pointer with a memory location it zeroes out that memory location to the default value of that type

	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value p points to is: %v \nthe value of i is %v", *p, i)
	// set the value that the memory location p is pointing to to 10
	*p = 10
	// this means that we want the memory address of the variable and not it value
	ptr := &i
	*ptr = 10

	fmt.Printf("The value p points to is: %v \nthe value of i is %v", *p, i)
}
