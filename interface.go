package main

import "fmt"

/**
1. interface adalah tipe data abtrack, dan tidak memiliki interface langsung
2. sebuah interface berisi definisi-definisi method
3. biasanya interface digunakan sebagai kontrak
*/

/**
 # implementasi interface
1. setiap tipe data yang sesuai dengan kontrak  interface, secara otomatis akan dianggap interface
2. sehingga kita tidak perlu menngimplementasikan interface secara manual
*/

// kontrak interface
type HasName interface {
	getName() string
}

func sayHello(hasname HasName) {
	fmt.Println("hallo nam saya adlah, ", hasname.getName())
}

//==============================================================

// implemantasi kedua interface

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

// ==================================================================

// implemantasi ketiga interface
type Animal struct {
	Name string
}

func (animal Animal) getName() string {
	return animal.Name
}

func main() {
	// car akses interface, implementasi kedua
	var orang Person
	orang.Name = "budai"
	sayHello(orang)

	// cara akses interface, implementasi ketiga
	cat := Animal{
		Name: "kucing",
	}
	sayHello(cat)

}
