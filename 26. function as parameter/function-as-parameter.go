package main

import "fmt"

// gunakan type parameter ketika parameter teralu panjang
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

// contoh lain

func mobilWithFilter(mobil string, filter func(string) string) {
	mobilFiltered := filter(mobil)
	fmt.Println("Kmu adalah mobil", mobilFiltered)
}

func mobilSpam(mobil string) string {
	if mobil == "honda" {
		return "..."
	} else {
		return mobil
	}
}

func main() {
	sayHelloWithFilter("budi", spamFilter)

	// atau disipan dalam variabel, cara memangil nya
	// filter := spamFilter
	// fmt.Println("anjing", filter)

	// contoh lain
	mobilWithFilter("honda", mobilSpam)
}
