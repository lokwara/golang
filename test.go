package main

import (
	"fmt"
)

func updateName(x *string) {
	*x = "wedge"
}

func main() {
	name := "Lewis"

	//updateName(name)

	//fmt.Println("memory address of name is:", &name)

	m := &name
	fmt.Println("memory address:", m)
	fmt.Println("value at memory address:", *m) //histeric gives you value at memory address, without gives you memory address

	fmt.Println(name)
	updateName(m)
	fmt.Println(name)

}
