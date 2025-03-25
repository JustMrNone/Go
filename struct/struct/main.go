// In summary, encapsulation is about hiding internal details and providing controlled access to them, while polymorphism is the ability to use entities of different
// types interchangeably based on their commonalities (usually through inheritance and method overriding).

package main

import (
	"fmt"
)

// structs can hold mixed types in the form of fields that we define by name
// fields can be anything that you want even a struct
type Engine struct {
	mpg      uint8
	gallon   uint8
	carOwner owner
}
type ElectricEngine struct {
	mpkwh uint8
	kwh   uint8
}
type owner struct {
	name string
}

// structs also have concepts of methods that we can use as well
// these are functions that are directly tied to the struct and have access to the strucst instance itself
// (e Engine) this assignes the function to the engine type
func (e Engine) MilesLeft() uint8 {
	return e.gallon * e.mpg
}

func (e ElectricEngine) MilesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	MilesLeft() uint8
}

func canMakeit(e engine, miles uint8) {
	if miles <= e.MilesLeft() {
		fmt.Println("You can make it, bitch")
	} else {
		fmt.Println("you are out of luck, hoe")
	}
}

func main() {
	//struct literal syntax
	var MyEngine Engine = Engine{mpg: 100, gallon: 200, carOwner: owner{"None"}}
	fmt.Println(MyEngine.mpg, MyEngine.gallon, MyEngine.carOwner.name)

	// you can also ommit the field names
	var MyEngine2 Engine = Engine{25, 40, owner{"MrHitch"}}

	//or

	MyEngine2.mpg, MyEngine2.gallon, MyEngine2.carOwner.name = 30, 20, "Jimbo"

	Electric := ElectricEngine{25, 60}

	fmt.Println(MyEngine2.mpg, MyEngine2.gallon, MyEngine2.carOwner.name)

	//we can also define anonymous struct but we need to define and use it at the same place

	var MyEngine3 = struct {
		mpg     uint8
		gallons uint8
	}{25, 100}
	// this is not reusable
	fmt.Println(MyEngine3.mpg, MyEngine3.gallons)
	Left := MyEngine2.MilesLeft()
	fmt.Println(Left)
	//this can take both engines the as long as they both have miles left method
	canMakeit(MyEngine, 50)
	canMakeit(Electric, 40)

}
