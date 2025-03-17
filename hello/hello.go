package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Farthamiul", "Bitch", "JackAss"}

	messages, errs := greetings.Hellos(names)

	if errs != nil {
		log.Fatal(errs)
	}

	// Request a greeting message.
	message, err := greetings.Hello("Farthamiul Bucklenuts")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(messages)
	fmt.Println(message)
}
