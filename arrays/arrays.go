package main

import (
	"fmt"
)

func main() {
	var intarr [3]int32 // 4bytes
	fmt.Println(intarr[0])
	fmt.Println(intarr[1:3])

	// we can see each memory location with &
	// compiler does not know all of the memory locations it just needs to know the first one and increment by four to get the rest

	fmt.Println(&intarr[0])
	fmt.Println(&intarr[1])
	fmt.Println(&intarr[2])
	//we can immediately initilaze the array using this syntax
	var intArr [3]int32 = [3]int32{1, 2, 3}
	// or

	intArray := [3]int32{1, 2, 3}
	// we can evem ommit 3

	intMyArray := [...]int32{1, 2, 3}

	fmt.Println(intArr[0], intArr[1], intArr[2])
	fmt.Println(intArray[0], intArray[1], intArray[2])
	fmt.Println(intMyArray[0], intMyArray[1], intMyArray[2])

}
