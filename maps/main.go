package main

import (
	"fmt"
)

func main() {
	fmt.Println(`map is a key value pair that you can look up a value by it's key`)
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)
	var myMap2 = map[string]uint{"Cutton Eye joe": 24, "Grease ball jim": 23}
	fmt.Println(myMap2["Cutton Eye joe"])
	fmt.Println(myMap2["Grease ball jim"])
	// if I Want to get a value from the list that does not exist in the list it returns the default value for that data type

	fmt.Println(myMap2["Monkey face jack"])

	// maps in go also return an optional boolean which returns true if the boolean is in the map

	var age, yes = myMap2["Cutton Eye joe"]
	fmt.Println(age, yes)

	var age2, yes2 = myMap2["Monkey face jack"]
	fmt.Println(age2, yes2)

	if yes {
		fmt.Printf("the age is %v", age)
	} else {
		fmt.Println("invalid name.")
	}
	if yes2 {
		fmt.Printf("the age is %v", age)
	} else {
		fmt.Println("\ninvalid name.")
	}

	// to delete an element from the map we use
	//delete(map, key)
	delete(myMap2, "Grease ball jim")

}
