package type_data

import "fmt"

func TypeNumber() string {
	return typeNumber()
}

//number tanpa argumen
func typeNumber() string {

	var title string = "type_data - number"
	fmt.Println(title)

	var num1 int
	var num2 int
	var num3 int

	num1 = 0
	num2 = 1
	num3 = num1 + num2

	fmt.Println("Number 1 = ", num1)
	fmt.Println("Number 2 = ", num2)
	fmt.Println("Number 3 = ", num3)
	return title
}
