package main

import (
	"fmt"
	"sort"
)

func main() {
	greeting := "Hey winner"

	// examples

	// fmt.Println(strings.Contains(greeting, "winner"))            // true, checks if the string contains the substring
	// fmt.Println(strings.ReplaceAll(greeting, "winner", "loser")) // Hey loser, replaces all the occurrences of the substring
	// fmt.Println(strings.ToUpper(greeting))                       // HEY WINNER, converts the string to uppercase
	// fmt.Println(strings.Index(greeting, "ey")) // 3, returns the index of the first occurrence of the substring
	// fmt.Println(strings.Split(greeting, " ")) // [Hey winner], splits the string into a slice of substrings

	// the original string is not modified
	fmt.Println("original string:", greeting) // Hey winner

	// sorting

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25, 34}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30) // if the value is not found, it returns the index where it should be inserted

	fmt.Println(index) // 2, the index of the first occurrence of 30 in the sorted slice

	// sorting with strings

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	sort.Strings(names)
	fmt.Println(names) // [bowser luigi mario peach yoshi]

	fmt.Println(sort.SearchStrings(names, "mario")) // 2, the index of the first occurrence of "mario" in the sorted slice
}
