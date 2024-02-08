package main

import "fmt"

// Structs are a way to create custom types, aka classes like in Python.

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// You can also create a method for a struct aka new bills, like this:
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}

	return b
}

// You can also create a method for a struct aka format bills, like this:
func (b bill) format() string {
	fs := "Bill breakdown:\n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-15v ...$%v \n", k+":", v) // -15v pushes the value to the right
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-15v ...$%0.2f \n", "tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-15v ...$%0.2f", "total:", total+b.tip)

	return fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	(*b).tip = tip
}

// add an item to the bill
func (b bill) addItem(name string, price float64) {
	b.items[name] = price
}

type Person struct {
	Name string
	Age  int
}

func structsvar() {
	p := Person{Name: "Tommy", Age: 25}

	fmt.Println(p.Name)
	fmt.Println(p.Age)
}
