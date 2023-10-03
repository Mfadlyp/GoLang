package main

import "fmt"

func getPerson() (string, string) {
	return "fadly", "pangestu"
}

func main() {
	firstName, lastName := getPerson()

	// cara menghiraukan return value _
	//firstName, _ := getPerson()
	fmt.Println("nama depan : ", firstName, "nama belakang : ", lastName)
}
