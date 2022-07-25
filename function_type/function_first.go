package function_type

import "fmt"

func FunctionFirst() string {
	var title string = "function_type - function_first"

	for i := 0; i < 10; i++ {
		fmt.Println("First Function Tutorial")
	}
	return title
}
