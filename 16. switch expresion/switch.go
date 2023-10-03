package main

import "fmt"

func main() {
	name := "jawir ireng"

	switch name {
	case "jawir":
		fmt.Println("Kamu adalah jawir")
	case "budi":
		fmt.Println("kamu adalah budi")
	default:
		fmt.Println("Tidak ada yang benar nama nya")
	}

	// switch short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama kamu lebih dari 5")
	case false:
		fmt.Println("Nama kamu kurang dari 5")
	}

	// switch tanpa kondisi
	panjang := len(name)
	switch {
	case panjang > 10:
		fmt.Println("nama kamu sangat panjang")
	case panjang > 5:
		fmt.Println("nama kamu sangat pendek")
	default:
		fmt.Println("Silakan cek kembali")
	}
}
