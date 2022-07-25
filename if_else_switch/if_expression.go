package if_else_switch

import "fmt"

func IfExpression() string {
	var title string = "if_else_switch - if"
	fmt.Println(title)

	var name string = "asoy"

	if name == "Gibran" {
		fmt.Println("Haloo Gibran")
	} else if name == "Utee" {
		fmt.Println("Haloo Utee")
	} else if name == "Gibs" {
		fmt.Println("Haloo Gibs")
	} else {
		fmt.Println("SIAPA LU? ", name)
	}

	// short statement
	// var length = len(name) //Normal statement
	fmt.Println("Deklarasi variable didalam if statement / if short statement")

	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}

	return title
}
