package main

import "fmt"

func main() {
	var name string

	name = "fadly"
	fmt.Println(name)

	name = "fadly pangestu" //merubah isi dari variabels
	fmt.Println(name)

	// membuat varibael secara langsung tanpa memangil tipe data lagi
	var name_2 = "fadly"
	fmt.Println(name_2)

	// boolean
	var hasil = true
	fmt.Println(hasil)

	// int
	var angka = 13
	fmt.Println(angka)

	// memaksa menggunakan int8
	var age int8 = 30
	fmt.Println(age)

	// cara tidak menggunakan var saat inisialisasi varibael :=s
	umur := 20
	fmt.Println(umur)

	// membuat varibel lebih dari satu
	var (
		firstname = "fadly"
		lastname  = "pangestu"
	)
	fmt.Println(firstname, lastname)
}
