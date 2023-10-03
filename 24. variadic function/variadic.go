package main

import "fmt"

// variadic function , ditandai dengan ... pada parameter
func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers { // masuk ke sini akan berupa slice
		total += number
	}
	return total
}

func main() {
	sumTotal := sumAll(10, 10, 10, 10, 10)
	fmt.Println(sumTotal)

	// slice parameter
	angka := []int{10, 20, 30}     // slice
	angkaTotal := sumAll(angka...) // tambahkan ...
	fmt.Println(angkaTotal)
}
