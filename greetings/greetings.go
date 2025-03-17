// Go code is grouped into packages, and packages are grouped into modules.
// Your module specifies dependencies needed to run your code,
// including the Go version and the set of other modules it requires.
package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

//In Go, a function whose name starts with a capital letter can be called by a function not in the same package.
// This is known in Go as an exported name.

//Go, the := operator is a shortcut for declaring and initializing a variable in one line the long way is

func Hello2(name string) string {
	var message string
	message = fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
