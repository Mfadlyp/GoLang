package main

import "fmt"

func main() {
	name := "bejo"

	if name == "fadly" {
		fmt.Println("hai, ", name)
	} else {
		fmt.Println("kamu bukan ", name)
	}

	// if short statement
	if length := len(name); length > 5 {
		fmt.Println("nama anda panjang")
	} else {
		fmt.Println("nama anda pendek")
	}
}
