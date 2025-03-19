package main

import (
	"errors"
	"fmt"
)

func main() {
	var num, deno int

	fmt.Print("Enter numerator: ")
	fmt.Scan(&num)

	fmt.Print("Enter denominator: ")
	fmt.Scan(&deno)

	var result, remainder, err = intDivRem(num, deno)
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("%v / %v = %v", num, deno, result)
	} else {
		fmt.Printf("%v / %v = %v, %v %% %v = %v", num, deno, result, num, deno, remainder)
	}

}

func intDivRem(numerator int, denomenator int) (int, int, error) {
	var err error

	if denomenator == 0 {
		err = errors.New("division by zero error")
		return 0, 0, err
	}

	division := numerator / denomenator
	reminder := numerator % denomenator

	return division, reminder, err
}
