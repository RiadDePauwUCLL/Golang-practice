package main

import "fmt"

var score = 100

func packagescope() {

	sayHi("mario")

	for _, v := range ages {
		fmt.Println(v)
	}

	showScore()
}
