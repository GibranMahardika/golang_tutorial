package type_data

import "fmt"

// Jalankan salah satu fungsi dengan comment dari salah satu fungsinya tersebut

func Random() interface{} {
	return true
}

// func TypeAssertionsImplementation() string {
// 	var title string = "type_data - type_assertions - TypeAssertionsImplementation"
// 	var result interface{} = Random()
// 	var resultString string = result.(string)
// 	fmt.Println(resultString)

// 	return title
// }

func TypeAssertionsSwitch() string {
	var result interface{} = Random()
	var title string = "type_data - type_assertions - TypeAssertionsSwitch"
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is String")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown Type")
	}
	return title
}
