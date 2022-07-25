package function_type

import "fmt"

func FunctionVariadic(numbers ...int) int {
	var title string = "function_type - function_variadic"
	fmt.Println(title)

	total := 0
	for _, value := range numbers {
		total += value
	}

	return total
}
