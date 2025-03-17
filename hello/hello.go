package main

import (
	"fmt"

	"github.com/JustMrNone/Go/greetings"
)

func main() {
	message := greetings.Hello("Bitch")
	fmt.Println(message)
}
