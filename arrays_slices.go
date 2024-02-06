package main

import "fmt"

func arrays_slices() {
	// var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30} // the [3] is the length of the array, non-mutable just like tuples in Python.

	names := [4]string{"Riad", "John", "Paul", "Tommy"}

	fmt.Println("\n", ages, len(ages))                                      // [20 25 30] 3 (length of the array)
	fmt.Println(names, "are the names and this is the length:", len(names)) // [Riad John Paul] 3 (length of the array)

	// slices (useful for working with arrays)
	var scores = []int{100, 50, 60} // no need to specify the length of the array, it's mutable so consider it as a list in Python.
	scores[2] = 25
	scores = append(scores, 85) // append a new value to the slice

	fmt.Println(scores, "is the list of scores & this is the length:", len(scores)) // [100 50 25 85] 4 (length of the slice)

	// slice ranges (similar to Python)
	rangeOne := names[1:3] // [John Paul]
	rangeTwo := names[2:]  // [Paul Tommy]

	fmt.Println(rangeOne, "is the 1st slice & this is the 2nd one:", rangeTwo)

	rangeOne = append(rangeOne, "Simon") // append a new value to the slice
	fmt.Println(rangeOne, "is the 1st slice after appending a new value")
}
