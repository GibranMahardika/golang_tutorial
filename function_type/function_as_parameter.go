package function_type

import "fmt"

// type wordFilter func(string) string

func FunctionAsParameter(nameParam string, filter func(string) string) string {
	title := "function_type - function_as_parameter"
	nameFiltered := filter(nameParam)
	fmt.Println("Hello ", nameFiltered)
	return title
}

func FunctionAsParameter1(nameParam string) string {
	if nameParam == "Anjing" {
		return "..."
	} else {
		return nameParam
	}
}
