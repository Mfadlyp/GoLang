package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	// result := random()              // value yg disimpan berupa interface kosong
	// resultString := result.(string) // menggunakan type assertions untuk merubah dari interface kosong ke string
	// fmt.Println(resultString)

	// type assertion yang lebih bagus
	result := random()

	switch value := result.(type) { // mengecek tipe data pada random()
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("Unkown", value)
	}
}
