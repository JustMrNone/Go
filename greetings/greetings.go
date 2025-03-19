// Go code is grouped into packages, and packages are grouped into modules.
// Your module specifies dependencies needed to run your code,
// including the Go version and the set of other modules it requires.
package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		err := errors.New("no name was given")
		return "", err
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	//Add nil (meaning no error) as a second value in the successful return.
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}
	return messages, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	//Note that randomFormat starts with a lowercase letter, making it accessible only to code in its own package (in other words, it's not exported).
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

//In Go, a function whose name starts with a capital letter can be called by a function not in the same package.
// This is known in Go as an exported name.

//Go, the := operator is a shortcut for declaring and initializing a variable in one line the long way is

/*
func Hello2(name string) string {
	var message string
	message = fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
*/
