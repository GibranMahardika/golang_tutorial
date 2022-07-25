package type_data

import "fmt"

func TypeBoolean() string {

	var title string = "type_data - boolean"
	fmt.Println(title)

	var benar bool = true
	var salah bool = false
	fmt.Println("Benar = ", benar)
	fmt.Println("Salah = ", salah)
	return title
}
