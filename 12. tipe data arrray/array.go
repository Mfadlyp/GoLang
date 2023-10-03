package main

import "fmt"

func main() {
	// cara membuat array
	var kumpulanNama [3]string
	kumpulanNama[0] = "m"
	kumpulanNama[1] = "fadly"
	kumpulanNama[2] = "pangestu"

	fmt.Println(kumpulanNama[0])
	fmt.Println(kumpulanNama[1])
	fmt.Println(kumpulanNama[2])

	// cara simple
	var values = [3]int{
		80,
		90,
		100,
	}
	fmt.Println(values)
	fmt.Println("Panjang array : ", len(values))

}
