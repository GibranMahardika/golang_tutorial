package package_golang

import (
	"fmt"
	"os"
)

func PackageOs() string {
	var title string = "package_golang - package_os"
	args := os.Args
	fmt.Println("Argument :")
	fmt.Println(args)

	// fmt.Println(args[1])
	// fmt.Println(args[2])
	// fmt.Println(args[3])

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname : ", name)
	} else {
		fmt.Println("Error : ", err.Error())
	}
	return title
}
