package main

import "fmt"

func main() {
	age := 22
	name := "Lewis"

	/*
	    imported fmt use,
	   1.Print function
	*/
	fmt.Println("Hello", name, "you are", age)

	//formatted string
	/*format specifiers (
	%v - default format for variables(strings,integers,boolean etc)
	%s - format for strings
	%d - format for integers
	%f - format for float
	%T - format for type of variable

	)
	*/
	fmt.Printf("My name is %v and my age is %v", name, age)
	fmt.Printf("age is of %T", age)

	//Sprintf
	/*

	 */
	var str = fmt.Sprintf("My age is %v and my name is %v\n", age, name)
	fmt.Println(str)
}
