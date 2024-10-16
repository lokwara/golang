package main

import "fmt"

func main() {
	// strings
	var nameOne string = "Lewis"
	//integer
	var age int = 22
	//float
	var deci float32 = 2.2
	var dec float64 = 2.5
	//boolean
	var yes bool = false

	// another way of declaring variables.
	// You dont have to give variable type and var
	// cant be used out of a function
	nameTwo := "Lewis Okoche"
	ages := 22
	decis := 3.142
	nooo := true

	fmt.Println(nameOne, age, deci, yes, dec, nameTwo, ages, decis, nooo)
}
