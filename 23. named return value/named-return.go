package main

import "fmt"

// named return value
func getName() (firstName string, middleName string, lastName string) {
	firstName = "m"
	middleName = "fadly"
	lastName = "pangestu"

	return //firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getName()

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
