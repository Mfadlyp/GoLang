package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address_1 := Address{"Subang", "Bandung", "Jakarta"}
	address_2 := &address_1

	address_2.City = "Jawa"
	fmt.Println("address_2", address_2)

	// addrress_2 mereubah isi value denngan asterisk operator *
	// menggunakan * akan merubah refrence dari yg lama ke yg baru (address_2)
	// sebelum address_1 mengacu ke Address subang
	// sesduah address_1 mengacu ke Address depok mengikutin address_2

	*address_2 = Address{"Depok", "Bekasi", "Bogor"}

	fmt.Println("address_1", address_1)
	fmt.Println("address_2", address_2)
}
