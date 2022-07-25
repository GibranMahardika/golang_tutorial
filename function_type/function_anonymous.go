package function_type

import "fmt"

type Blacklist func(string) bool

// pengertian anonymous function yaitu, membuat function tanpa menggunakan nama functionnya
// function ini digunakan dengan variable, biasanya dilakukan pada main.

func FunctionAnonymous(nameBlacklist string, blacklist Blacklist) string {
	if blacklist(nameBlacklist) {
		fmt.Println("You are blocked", nameBlacklist)
	} else {
		fmt.Println("Welcome", nameBlacklist)
	}
	return nameBlacklist

}

// func AnonymousBlacklistAdmin(nameAdmin string) bool {
// 	return nameAdmin == "admin"
// }

// func AnonymousBlacklistRoot(nameRoot string) bool {
// 	return nameRoot == "root"
// }
