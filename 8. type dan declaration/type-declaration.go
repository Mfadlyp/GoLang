package main

import "fmt"

func main() {
	// type membuat alias
	// membuat alias string dengan noKtp
	type noKtp string

	var nomorktp noKtp = "12ed43"

	fmt.Println(nomorktp)

	// contoh lain
	type maried bool
	var mariedStatus maried = true
	fmt.Println(mariedStatus)
}
