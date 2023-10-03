package main

import "fmt"

// recursive function adalah function yang memangil diri nya sendiri

// function factorial angka
func loopAngka(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

// rescursive function
func recursiveFunction(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * recursiveFunction(value-1) // memangil function itu sendiri
	}
}

func main() {
	fmt.Println(loopAngka(5))

	// memangil function recursive
	re := recursiveFunction(5 )
	fmt.Println(re)
}
