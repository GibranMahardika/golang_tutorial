package main

import "fmt"

func main() {
	type noKTP string
	type married bool

	var ktpGibs noKTP = "0123456789"
	var marriedStatus married = false
	fmt.Println(ktpGibs)
	fmt.Println()
	fmt.Println(marriedStatus)
	fmt.Println()
	fmt.Println(noKTP(""))
}
