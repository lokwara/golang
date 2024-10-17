package main

import (
	"fmt"
)

// Go is a Pass-by-value Language

func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	// group A types -> strings, ints, bools, arrays, structs
	name := "Lewis"

	name = updateName(name)

	fmt.Println(name)

	//group B types -> slices, maps, functions
	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}

	updateMenu(menu)
	fmt.Println(menu)

}
