package main

import (
	"fmt"
)

func main() {
	var myArr0 [4]int = [4]int{10, 200, 30, 40}
	myArr := [...]int{10, 30, 20, 40, 60}

	for j := range myArr0 {
		fmt.Println(myArr0[j], ":", &myArr0[j])
	}

	for i := range myArr {
		fmt.Println(&myArr[i], ":", myArr[i])
	}
	fmt.Println(len(myArr), len(myArr0))

}
