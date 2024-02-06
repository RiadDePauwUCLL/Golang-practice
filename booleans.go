package main

import "fmt"

func booleans() {

	// Booleans, if, else, else if
	age := 25

	// basic comparison, true or false
	// fmt.Println(age <= 50)
	// fmt.Println(age >= 50)
	// fmt.Println(age == 45)
	// fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("Age is less than 30")
	} else if age < 40 {
		fmt.Println("Age is less than 40")
	} else {
		fmt.Println("Age is greater than 40")
	}

	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("Continuing at pos", index)
			continue // skip the rest of the code in the loop for this iteration
		}
		if index > 2 {
			fmt.Println("Breaking at pos", index)
			break // quit the loop completely
		}
		fmt.Printf("The value at pos %v is %v\n", index, value)
	}
}
