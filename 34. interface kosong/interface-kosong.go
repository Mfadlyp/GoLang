package main

import "fmt"

// contoh lain dari interface kosong
func kosong() any {
	// return true
	// reeturn "salah"
	return 1
}

func ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		fmt.Println("ups salah")
	}
	return i
}

func main() {
	// kosong := ups(1)
	// fmt.Println(kosong)

	var isi any = kosong()
	fmt.Println(isi)

	var data interface{} = ups(1)
	fmt.Println(data)
}
