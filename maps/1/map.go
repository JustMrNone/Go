package main

import (
	"fmt"
)

func main() {
	myMap := map[string]int{
		"apple":  1,
		"banana": 2,
		"Milk":   3,
	}
	fmt.Println(myMap["apple"])
	_, ok := myMap["orange"]
	if !ok {
		fmt.Println("value not found!")
	}

	myMap["Grapes"] = 4

	for key, value := range myMap {
		fmt.Print("\n")
		fmt.Printf("key: %v, value: %v", key, value)
	}

	keys := []string{"apple", "banana", "Milk", "Grapes"}

	for _, key := range keys {
		fmt.Printf("\nkey: %v, value: %v", key, myMap[key])
	}
}
