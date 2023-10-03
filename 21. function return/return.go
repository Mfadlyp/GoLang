package main

import "fmt"

func getHello(nama string) string {
	return "hello" + nama
}

func main() {
	result := getHello("budi")
	fmt.Println(result)
}
