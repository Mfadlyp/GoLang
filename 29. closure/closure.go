package main

import "fmt"

func main() {
	name := "budi"
	counter := 0
	increament := func() {
		name = "eco" // akan mengubah isi dari variabel name menjadi eco dari budi
		fmt.Println("data increament")
		counter++ // akan menambah counter diluar func variabel ini closure
	}
	increament()
	increament()
	fmt.Println(counter)
	fmt.Println(name)
}
