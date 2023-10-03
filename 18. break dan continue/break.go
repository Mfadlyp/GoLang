package main

import "fmt"

func main() {
	// menggunakan break
	for i := 0; i < 10; i++ {
		if i == 5 {
			// fmt.Println("lima ditemukan")
			break
		}
		fmt.Println("Perulangan yang ke-", i)
	}

	// menggunakan continue
	for angka := 0; angka < 10; angka++ {
		if angka%2 == 0 {
			continue
		}
		fmt.Println("Perulangan yang ke-", angka)
	}
}
