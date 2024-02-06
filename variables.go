package main

import "fmt"

var someName = "gaming"

// strings & what not
func variables() {
	var nameOne string = "Mario"
	var nameTwo string = "Luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Bowser"
	nameThree = "Peach"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Yoshi"
	fmt.Println(nameFour)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory. Always use int unless you need to use a specific size, look documentation for more info

	// var numOne int8 = 25      // int8 means 8 bits
	// var numTwo int16 = -129   // int16 means 16 bits
	// var numThree uint16 = 256 // unsigned int, only positive numbers

	// floats, 32 & 64 bits. Always use float64 unless you need to use a specific size, look documentation for more info

	// var scoreOne float32 = 25.98
	// var scoreTwo float64 = 25.987654321
	// scoreThree := 1.5 // float64, default.
}
