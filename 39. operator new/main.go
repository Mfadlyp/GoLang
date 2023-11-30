package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// kata kunci new untuk membuat pointer baru
	address1 := new(Address)
	address2 := address1

	address2.City = "Indonesia"

	fmt.Println(address1)
	fmt.Println(address2)
}
