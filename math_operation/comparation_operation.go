package math_operation

import "fmt"

func CompOperation() string {

	var title string = "math_operation - comparation_operation"
	fmt.Println(title)

	var nama string = "Gibs"
	var nama1 string = "Gibs"
	var nama2 = "Gibran"

	var compare = nama == nama1
	var compare1 = nama1 == nama2
	var compare2 = nama > nama1
	var compare3 = nama1 > nama2

	var num = 10
	var num1 = 100

	var compareNum = num >= num1
	var compareNum1 = num <= num1
	var compareNum2 = num == num1
	var compareNum3 = num != num1

	fmt.Println(compare)
	fmt.Println(compare1)
	fmt.Println(compare2)
	fmt.Println(compare3)
	fmt.Println()

	fmt.Println(compareNum)
	fmt.Println(compareNum1)
	fmt.Println(compareNum2)
	fmt.Println(compareNum3)
	fmt.Println()

	fmt.Println(num >= num1)
	fmt.Println(num <= num1)
	fmt.Println(num == num1)
	fmt.Println(num != num1)
	return title
}
