package main

import "fmt"

func maps() {

	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55, // ending with a comma is mandatory
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// Looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		243434:     "mario",
		8934859435: "luigi",
		1203921093: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[243434])

	// In case you want to find a key by value
	nameToFind := "mario" // The name you're looking for

	for key, value := range phonebook {
		if value == nameToFind {
			fmt.Println(key)
			break
		}
	}

	// changing a value
	phonebook[243434] = "bowser"
	fmt.Println(phonebook)
}
