package main

import "fmt"

func loops() {

	// basic for loop

	// x := 0
	// for x < 5 {
	// 	fmt.Println("Value of x is", x)
	// 	x++
	// }

	// more compact for loop

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Value of i is", i)
	// }

	names := []string{"mario", "luigi", "yoshi", "peach"}

	// prints out each name in the names slice
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// range keyword is used to loop over an array, slice, string, map or channel
	for index, value := range names {
		fmt.Printf("The value at index %v is %v\n", index, value)
	}

	// if you don't need the index, you can use the _ character. This is called a blank identifier, and is used when you want to ignore a value.
	for _, value := range names {
		fmt.Printf("The value is %v\n", value)
		value = "new string" // this will not change the value in the names slice, as value is a copy of the value in the slice
	}

	fmt.Println(names)
}
