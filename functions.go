package main

import (
	"fmt"
	"math"
)

func sayHello(n string) {
	fmt.Printf("Hello %v!\n", n)
}

func sayBye(n string) {
	fmt.Printf("Bye %v!\n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 { // Function that returns a float64, takes a float64 as an argument
	return math.Pi * r * r
}

func functions() {

	// Call the functions
	// sayHello("mario")
	// sayHello("luigi")
	// sayBye("mario")
	// sayBye("luigi")

	// Pass a function as an argument
	cycleNames([]string{"mario", "luigi", "yoshi"}, sayHello)
	cycleNames([]string{"mario", "luigi", "yoshi"}, sayBye)

	// Call the circleArea function
	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("Circle 1 is %.2f and Circle 2 is %.2f\n", a1, a2) // formatting similar to Python
}
