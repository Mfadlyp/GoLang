package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr . " + man.Name
}

func main() {
	fadly := Man{"fadly"}
	fadly.Married()

	fmt.Println(fadly.Name)
}
