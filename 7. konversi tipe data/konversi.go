package main

import "fmt"

func main() {
	// konversi 32 ke 64
	var nilai32 int32 = 3245
	var nilai64 int64 = int64(nilai32)

	fmt.Println(nilai64)

	// koversi byte ke string
	var name = "fadly"
	var x byte = name[0]

	var nameStr string = string(x)

	fmt.Println(nameStr)

}
