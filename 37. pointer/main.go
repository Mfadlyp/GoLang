package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// passing by value dengan meng copy value yg ada
	address_1 := Address{"bekasi", "Jakarta", "Indonesia"}
	address_2 := address_1  // passing by value
	address_3 := &address_1 // pointer  > pass by refrence

	address_2.City = "bandung"
	address_3.City = "Jawa"

	fmt.Println(address_1) // address_1 tidak akan berubah
	fmt.Println(address_2)
	fmt.Println(address_3) // pointer mengarah ke address_1
	fmt.Println("Setelahhh merubah", address_1)

	// pointer adalah pass by refrence
}
