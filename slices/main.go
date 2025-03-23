package main

import (
	"fmt"
)

func main() {
	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	//with slices we can add to the slice using the built in append function
	//it takes the slice and the value that we want to add
	//then it will return the newslice that have the new element
	fmt.Printf("length = %v | capacity = %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("length = %v | capacity = %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

	//we can add multiple values to our slice using the spread operator like this

	var intMySlice []int32 = []int32{8, 9}

	intSlice = append(intSlice, intMySlice...)

	fmt.Println(intSlice)

	// another way to create a slice is to use the make function

	intSlice3 := make([]int, 3)

	for i := range len(intSlice3) {
		intSlice3[i] = i
	}
	fmt.Printf("\n")
	fmt.Println(intSlice3)
	fmt.Println(cap(intSlice3), len(intSlice3))

	intSlice3 = append(intSlice3, 3, 4, 5)
	fmt.Println(intSlice3)

	fmt.Println(cap(intSlice3), len(intSlice3))
}
