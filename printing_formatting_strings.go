package main

import "fmt"

func formatting_strings() {

	name := "Riad"
	number := 10

	// Print
	// fmt.Print("Hello, ")
	// fmt.Print("World! \n")

	// Println
	fmt.Println("Hi y'all!")
	fmt.Println("My name is", name, "and I like", number, "chicken nuggets!")

	// Printf (formatted strings)
	fmt.Printf("My name is %v and I like %v chicken nuggets! \n", name, number) // %v is a placeholder for the value
	fmt.Printf("My name is %q and I like %q chicken nuggets! \n", name, number) // %q adds double quotes around the string
	fmt.Printf("My name is %T and I like %T chicken nuggets! \n", name, number) // %T prints the type of the value
	fmt.Printf("I like %f chicken nuggets! \n", 5.50)                           // %f prints the float value

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("My name is %v and I like motorcycles! \n", name) // %v is a placeholder for the value
	fmt.Println("The saved string is:", str)
}
