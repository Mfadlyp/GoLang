package main

import "fmt"

func main() {
	// versi singkat
	person := map[string]string{
		"nama": "fadly",
		"age":  "20",
	}
	// cara merubah/menambah data nya
	person["title"] = "Yang kocag"

	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["age"])

	// versi panjang
	var orang map[string]string = map[string]string{
		"nama": "beji",
		"age":  "30",
	}
	fmt.Println(orang)

	// membuat map menggunakan make
	book := make(map[string]string)
	book["judul"] = "turu"
	book["nama"] = "budi"
	book["tahun"] = "2021"
	fmt.Println("sebelum delete = ", book)

	// menghapus map delete(map,key)
	delete(book, "tahun")
	fmt.Println("sesudah delete = ", book)
}
