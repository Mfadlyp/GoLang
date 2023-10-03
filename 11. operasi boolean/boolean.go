package main

import "fmt"

func main() {
	// operasi boolean
	var nilai_1 int = 85
	var nilai_2 int = 85

	var nilaiAwal bool = nilai_1 > 80
	var nilaiAkhir bool = nilai_2 > 80

	var hasil_and bool = nilaiAwal && nilaiAkhir
	var hasil_or bool = nilaiAwal || nilaiAkhir
	var hasil_not bool = nilaiAwal != nilaiAkhir

	fmt.Println(hasil_and)
	fmt.Println(hasil_or)
	fmt.Println(hasil_not)

	// yang simple
	fmt.Println("Yang lebih simple", nilai_1 >= 80 && nilai_2 >= 80)

}
