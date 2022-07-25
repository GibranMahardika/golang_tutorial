package type_data

import "fmt"

func TypeDeclaration() string {

	var title string = "type_data - type_declaration"
	fmt.Println(title)

	type noKtp int
	type marry bool

	var NoKTP noKtp = 123456789
	var married marry = false

	fmt.Println(NoKTP)
	fmt.Println(married)

	return title
}
