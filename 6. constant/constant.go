package main

import "fmt"

func main() {
	// const tidak rewel ketika variabel tidak digunakan
	// menggunakan const isi varibel tidak dapat dirubah
	const fisrtName string = "fadly"
	const lastName = "pangestu"

	fmt.Println(fisrtName, lastName)

	// multiple const
	const (
		nama string = "fadly"
		age         = 20
	)
	fmt.Println(nama, age)
}
