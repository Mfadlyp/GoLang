package main

import "fmt"

func main() {
	counter_1 := 1

	// kondisi bernilai true
	for counter_1 <= 10 {
		fmt.Println("Perulangan ke- ", counter_1)
		counter_1++
	}

	// for dengan statement

	// for init statement; value; post statement
	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println("New ulang ke- ", counter)
	// }

	// interasi menggunakan for
	slice := []string{"budi", "jawir", "eko"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// cara cepat mengambil data dari array, slice, map menggunkan for range
	name := []string{"eko", "budi", "jawir", "herman", "bejo"}

	// jikalau hanya mengambil nilai ny saja tanpa index
	// gunakan _
	for _, name := range name {
		// fmt.Println("Indes- ", index, "Nama- ", name)
		fmt.Println("nama nya adalah -", name) // hanya mengambil value saja
	}

}
