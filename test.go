package main

import "fmt"

func main() {
	/*
		Arrays
		len is length
		start by variable, variable name, then length of the array, followed by variable type
	*/
	var ages [3]int = [3]int{25, 57, 58}

	names := [4]string{"Silus", "Gibson", "Lewis", "Lebron"}
	names[1] = "gina" //this replaces the number 1 value from "Gibson" to "Gina"

	fmt.Println(ages, len(ages)) //the len prints length of array
	fmt.Println(names, len(names))

	//Slices
	/*
		we dont enter number in square brackets
	*/
	var scores = []int{100, 50, 60} //this slice makes it easy to extend number of items on array without limit
	scores[2] = 25
	scores = append(scores, 85) //appends element to end of slice, making it remove 60 and add 85

	fmt.Println(scores, len(scores))

	//slice ranges
	rangeOne := names[1:3]  //get me from position 1 to 3
	rangeTwo := names[2:]   //get me from position 2 to 3
	rangeThree := names[:3] //get upto position 3

	fmt.Println(rangeOne, rangeTwo, rangeThree)
	rangeOne = append(rangeOne, "Beatrice")
	fmt.Println(rangeOne)
}
