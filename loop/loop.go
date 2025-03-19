package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	var n int = 0
	for i := 0; i < 100000000; i++ {
		n += 1
	}

	elapsed := time.Since(start)
	fmt.Printf("Loop took %s\n", elapsed)
	//this is 120x faster than python
}
