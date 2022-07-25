package math_operation

import "fmt"

func Math() string {

	var title string = "math_operation - math_operation"
	fmt.Println(title)

	var number1 int = 1
	var number2 int = 2
	var number3 int = number1 + number2
	fmt.Println("Jumlah = ", number3)

	var i = 10
	i += 10 //Augmented Asignment
	fmt.Println("i = ", i)

	i = i + i
	fmt.Println("i Lanjutan = ", i)

	i++ //Unary Operation
	fmt.Println("i + 1 = ", i)

	i = -10
	fmt.Println("i minus = ", i)

	var positive = +100
	var negative = -100

	fmt.Println(positive)
	fmt.Println(negative)

	return title
}
