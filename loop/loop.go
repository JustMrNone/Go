package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Basic for loop with counter
	// This is the most common type of loop in Go
	start := time.Now()
	var n int = 0
	for i := 0; i < 100000000; i++ {
		n += 1
	}
	elapsed := time.Since(start)
	fmt.Printf("Basic loop took %s\n", elapsed)
	//this is 120x faster than python

	// 2. Range-based loop with slice
	// 'range' keyword provides both index and value
	var mySlice []int
	for i := range 10 {
		mySlice = append(mySlice, i)
	}

	// 3. Range loop with index only
	// When you only need the index, you can omit the value
	fmt.Println("\nLooping with index only:")
	for i := range mySlice {
		fmt.Println(mySlice[i])
	}

	// 4. Range loop with both index and value
	// This is the most idiomatic way to iterate over a slice
	fmt.Println("\nLooping with both index and value:")
	for i, v := range mySlice {
		fmt.Println(i, v)
	}

	// 5. Infinite loop with break
	// Demonstrates how to create an infinite loop with a break condition
	fmt.Println("\nInfinite loop with break:")
	counter := 0
	for {
		if counter >= 5 {
			break
		}
		fmt.Println("Count:", counter)
		counter++
	}

	// 6. While-like loop
	// Go doesn't have a 'while' keyword, but we can simulate it
	fmt.Println("\nWhile-like loop:")
	whileCounter := 0
	for whileCounter < 5 {
		fmt.Println("While count:", whileCounter)
		whileCounter++
	}

	// 7. Loop with continue
	// Demonstrates skipping iterations
	fmt.Println("\nLoop with continue:")
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // Skip this iteration
		}
		fmt.Println("Continue example:", i)
	}

	// 8. Nested loops
	fmt.Println("\nNested loops:")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}

	// 9. Range over string
	// Demonstrates how to iterate over a string's runes
	fmt.Println("\nRange over string:")
	for i, r := range "Hello" {
		fmt.Printf("Index: %d, Rune: %c\n", i, r)
	}

	// 10. Range over map
	// Shows how to iterate over a map
	fmt.Println("\nRange over map:")
	myMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
