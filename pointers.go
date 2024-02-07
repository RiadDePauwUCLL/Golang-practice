package main

import "fmt"

func updatePointer(x *string) {
	*x = "rambo"
}

func pointers() {
	name := "tifa"

	// updatePointer(name)

	// fmt.Println("Memory address of name is:", &name)

	m := &name // "&" gives you the memory address of the value
	// fmt.Println("Memory address of m is:", m)        // m is a pointer to the memory address of name
	// fmt.Println("Value at memory address m is:", *m) // *m is the value at the memory address m

	fmt.Println(name)
	updatePointer(m) // passing the memory address of name to the function
	fmt.Println(name)
}
