package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		Standard Library
	*/
	//greeting := "Hello Friends"

	//fmt.Println(strings.Contains(greeting, "hello"))
	//fmt.Println(strings.ReplaceAll(greeting, "Hello", "hi")) //subtituted the hello with hi on output

	//fmt.Println(strings.Index(greeting, "ll"))
	//fmt.Println(strings.Split(greeting, " "))
	//fmt.Println(strings.ToUpper(greeting))

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages) //method 'sort' sorts the integers from small to big
	fmt.Println(ages)

	index := sort.SearchInts(ages, 75) //searches and gives you where it is located
	fmt.Println(index)

	names := []string{"Lewis", "Lebron", "Beatrice", "Gibson", "Silus"}

	sort.Strings(names) //this method sorted names according to name and letter
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "Beatrice"))
}
