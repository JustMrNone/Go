package main

import (
	"fmt"
	"strings"
)

// go uses utf-8 encoding to repersent strings in your computer
// utf-32 uses 32 bits but
func main() {
	var myString string = "Résumé"
	var myString2 = []rune("Résumé")

	indexed := myString[0]
	indexed2 := myString2[1]

	fmt.Printf("%v, %T", indexed, indexed)
	fmt.Printf("%v %T", indexed2, indexed2)

	for index, value := range myString {
		fmt.Println(index, value)
	}
	for index, value := range myString2 {
		fmt.Println(index, value)
	}

	//we can declare a rune type like this

	myRune := 'a'

	fmt.Printf("%T", myRune)

	//we can concatinate using +

	mySlice := []string{"H", "i", " ", ",", "B", "I", "T", "C", "H"}
	var myStrings string = ""

	for i := range mySlice {
		myStrings += mySlice[i]
	}
	//strings are immutable in go
	//myStrings[0] = S ❌
	fmt.Println("\n", myStrings)

	mySlice2 := []string{"H", "i", " ", ",", "B", "I", "T", "C", "H"}

	var strBuilder strings.Builder

	for i := range mySlice2 {
		strBuilder.WriteString(mySlice2[i])
	}

	var catStr = strBuilder.String()
	fmt.Printf("%v", catStr)

}
