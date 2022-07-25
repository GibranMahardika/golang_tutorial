package if_else_switch

import "fmt"

func SwitchExpression() string {

	var title string = "if_else_switch - switch"
	fmt.Println(title)

	var name string = "Gibs"

	// Switch Case Default Statement
	fmt.Println("Switch Statement")

	switch name {
	case "Gibran":
		fmt.Println("Haloo Gibran")
	case "Gibs":
		fmt.Println("Haloo Gibs")
	case "Utee":
		fmt.Println("Haloo Utee")
	default:
		fmt.Println("SIAPA LU?", name)
	}

	// Switch Case Short Statement
	fmt.Println("Switch Short Statement")

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	// Switch Non Statement
	fmt.Println("Switch NON Statement")

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang", name)
	case length > 5:
		fmt.Println("Nama masih panjang", name)
	default:
		fmt.Println("Nama sudah mantap", name)
	}
	return title
}
