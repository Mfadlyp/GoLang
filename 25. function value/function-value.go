package main

import "fmt"

// function value adalah menyimpan function kedalam varibael
func getGoodBye(name string) string {
	return "hello " + name
}

func main() {
	goodbay := getGoodBye // simpan function getGoodBye kedalam variabel goobayy

	result := goodbay("fadly")
	fmt.Println(result)
	//fmt.Println(goodbay("eko")) // memangil functin value nya
}
