package main

import "fmt"

type blackList func(string) bool

func registUser(name string, Blacklist blackList) {
	if Blacklist(name) {
		fmt.Println("you are blcoked", name)
	} else {
		fmt.Println("wellcome", name)
	}
}

// jangan gunakan ini
// func blacklistAdmin(name string) bool {
// 	return name == "admin"
// }

// func blacklistUser(name string) bool {
// 	return name == "ujang"
// }

func main() {
	//           VV anonymous function
	Blacklist := func(name string) bool {
		return name == "admin"
	}
	registUser("admin", Blacklist)
	registUser("budi", Blacklist)

	// atau
	fmt.Println("admin", func(name string) bool {
		return name == "admin"
	})
}
