package main

import "fmt"

//variables can be used in other files
// to run all files just list them like this
//go run test.go greetings.go

var points = []int{20, 90, 100, 45, 70}

func sayHello(n string) {
	fmt.Println("hello", n)
}

func showScore() {
	fmt.Println("You scored this many points:", score)
}
