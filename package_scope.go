package main

import "fmt"

var score = 100

// package_scope2 is a package level variable, it can be accessed from any file in the package

func packagescope() {

	sayHi("mario")

	for _, v := range ages {
		fmt.Println(v)
	}

	showScore()
}
