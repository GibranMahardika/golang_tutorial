package function_type

import "fmt"

func FunctionParameter(firstName string, lastName string) string {

	var title string = "function_type - function_parameter"

	mixName := firstName + lastName

	fmt.Println("Hello ", mixName)
	fmt.Println("Hello ", firstName)
	fmt.Println("Hello ", lastName)
	fmt.Println("Hello ", mixName)
	return title
}
