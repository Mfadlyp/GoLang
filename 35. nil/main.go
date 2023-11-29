package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"Name": name,
		}
	}
}

func main() {
	data := newMap("fadly")
	if data == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(data)
	}
}
