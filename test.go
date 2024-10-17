package main

import (
	"fmt"
)

// Maps
func main() {

	menu := map[string]float64{
		"soup":      4.99,
		"Ugali":     2.99,
		"Sarat":     1.99,
		"Misheveve": 0.99,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])
	//Maps entire menu in the map
	//start with  map variable variable

	//looping maps
	//k,key is "soup", v,value is "4.99"
	for k, v := range menu {

		fmt.Println(k, "_", v)
	}

	//int as key type
	phonebook := map[int]string{
		769731531: "Lewis",
		710211186: "Mum",
		725797511: "Dad",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[769731531])

	phonebook[725797511] = "Silus"
	fmt.Println(phonebook)

	phonebook[710211186] = "Beatrice"
	fmt.Println(phonebook)

}
