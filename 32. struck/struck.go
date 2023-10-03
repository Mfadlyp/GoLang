package main

import "fmt"

/**
struck adalah sebuah data template
data distruck disimpan dalam field
*/

type Orang struct {
	Name, Addres string
	Age          int
}

// struck method
func (orang Orang) sayHello(name string) {
	fmt.Println("hai", name, "dan untuk kamu", orang.Name)
}

func main() {

	var orang Orang
	orang.Name = "Budi"
	orang.Addres = "jl.kocombrang"
	orang.Age = 20

	// cara akses struct methodss
	orang.sayHello("budai")

	// fmt.Println(orang.Name)

	// // struct literal
	// dataPerson := Orang{
	// 	Name:   "joko",
	// 	Addres: "jl. budi",
	// 	Age:    25,
	// }
	// fmt.Println(dataPerson.Name)
}
