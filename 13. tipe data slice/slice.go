package main

import "fmt"

func main() {
	var months = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"july",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	// membuat slice low ke high
	// slice dari low sampai sebelum high
	var slice1 = months[10:]
	fmt.Println("nilai slice 1 awalan = ", slice1)
	// fmt.Println("hasil slice 1 = ", slice1)
	// fmt.Println("panjang slice [5:] : ", len(slice1))   // untuk melihat panjang slice
	// fmt.Println("kapasitas slice [5:] : ", cap(slice1)) // untuk melihat kapasitas slice

	// fmt.Println("============================================================================")
	// merubah slice 0 / dari 4 => [4:8]
	// hati-hati merubah slice atau array akan saling merubah satu sama lain
	// slice1[0] = "mei lagi" // merubah slice 0 yang dari slice 1 bukan dari nilai array months
	// fmt.Println(slice1)

	// menambah slice baru menggunakan append
	newSlice := append(slice1, "menambahNilaiAppend")
	fmt.Println("sesudah append = ", newSlice)

	newSlice[0] = "ups"
	fmt.Println("append kembali = ", newSlice)

	fmt.Println("slice 1 : ", slice1)
	fmt.Println("Array Months = ", months)

	// 	fmt.Println("============================================================================")
	// 	// slice dari low sampai terakhir slice
	// 	var slice2 = months[4:]
	// 	fmt.Println("hasil slice 2 = ", slice2)
	// 	fmt.Println("panjang slice [4:] : ", len(slice2))   // untuk melihat panjang slice
	// 	fmt.Println("kapasitas slice [4:] : ", cap(slice2)) // untuk melihat kapasitas slice

	// 	fmt.Println("============================================================================")
	// 	// slice dari awal sampai sebelum akhir high
	// 	var slice3 = months[:6]
	// 	fmt.Println("hasil slice 3 = ", slice3)
	// 	fmt.Println("panjang slice [:6] : ", len(slice3))   // untuk melihat panjang slice
	// 	fmt.Println("kapasitas slice [:6] : ", cap(slice3)) // untuk melihat kapasitas slice

	// 	fmt.Println("============================================================================")
	// 	var slice4 = months[:]
	// 	fmt.Println("hasil slice 4 = ", slice4)
	// 	fmt.Println("panjang slice [:] : ", len(slice4))   // untuk melihat panjang slice
	// 	fmt.Println("kapasitas slice [:] : ", cap(slice4)) // untuk melihat kapasitas slice

	fmt.Println("============================================================================")
	// membuat slice dengan aman menggunakan make
	// make([]tipedata,low,high)
	var sliceBaru = make([]string, 2, 6)
	sliceBaru[0] = "data_1"
	sliceBaru[1] = "data_2"
	fmt.Println("slice baru make() = ", sliceBaru)

	fmt.Println("============================================================================")
	// mencopy slice
	baruSlice := make([]string, len(sliceBaru), cap(sliceBaru))
	// copy to, from
	copy(baruSlice, sliceBaru)
	fmt.Println("hasil copy() = ", baruSlice)

	// perbedaan array dengan slice
	iniArray := [4]int{0, 1, 2, 3}
	iniSlice := []int{0, 1, 2, 3}

	fmt.Println("ini array", iniArray)
	fmt.Println("ini array", iniSlice)
}
