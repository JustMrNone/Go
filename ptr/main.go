package main

import (
	"fmt"
)

func main() {

	Myint := 100
	Mystr := "Hello Fuckers"

	var myPtr *string = &Mystr
	MyPtr := &Myint

	MyString := *myPtr
	MyVar := *MyPtr
	MyPtrToPtr := &MyPtr

	*MyPtr = 20

	fmt.Println(**MyPtrToPtr)
	fmt.Println(MyPtr)
	fmt.Println(MyVar)
	fmt.Println(MyString)

}
