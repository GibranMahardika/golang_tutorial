package type_data

import "fmt"

func TypeString() string {

	var title string = "type_data - string"
	fmt.Println(title)

	var firstName string = "Gibran"
	var lastName string = "Mahardika"
	var nickName string = "Gibs"

	fmt.Println("First Name = ", firstName)
	fmt.Println("Last Name = ", lastName)
	fmt.Println("Nick Name = ", nickName)

	fmt.Println("Length First Name = ", len(firstName))
	fmt.Println("Catch character 0 = ", lastName[0])
	fmt.Println("Catch character 1 = ", nickName[1])

	return title
}
