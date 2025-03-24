package main

import (
	"fmt"
)

func main() {
	numbers := make([]int, 0, 6)
	numbers = append(numbers, 10, 15, 20, 25, 30, 35)

	for index, value := range numbers[2:] {
		fmt.Println(index, ":", value)
	}

	numbers2 := []int{10, 20, 50, 80, 100}

	for i := 2; i < len(numbers2); i++ {
		fmt.Println(i, ":", numbers2[i])
	}

}
