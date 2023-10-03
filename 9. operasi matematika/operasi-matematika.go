package main

import "fmt"

func main() {

	// operasi matematika

	// operasi penjumlahan
	angka_1 := 20
	angka_2 := 10
	hasil := angka_1 + angka_2
	fmt.Println("Hasil penjumlahan : ", hasil)

	// operasi pengurangan
	var angka_kurang1 int = 20
	var angka_kurang2 int = 10
	var hasil_kurang = angka_kurang1 - angka_kurang2
	fmt.Println("Hasil pengurangan : ", hasil_kurang)

	// operasi perkalian
	angka_kali1 := 10
	angka_kali2 := 10
	hasil_kali := angka_kali1 * angka_kali2
	fmt.Println("Hasil perkalian : ", hasil_kali)

	// operasi pembagian
	angka_bagi1 := 20
	angka_bagi2 := 5
	hasil_bagi := angka_bagi1 / angka_bagi2
	fmt.Println("Hasil pembagian : ", hasil_bagi)

	// operasi modulus
	angka_mod := 20 % 3
	fmt.Println("Hasil modulus : ", angka_mod)

	// aughmented assigments
	var i = 10
	i += 10        // i = i+10
	fmt.Println(i) // 20

	angka := 0
	angka += 10
	fmt.Println("Hasil augh : ", angka)

	angka++
	fmt.Println("hasil ++ : ", angka)

}
