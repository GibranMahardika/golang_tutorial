package package_golang

import (
	"fmt"
	"strconv"
)

func PackageStrconv() string {
	boolean, err := strconv.ParseBool("false")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(1000000, 2)
	//note 10 (decimal), 2 (binary), 8 (octal), 16(hexadecimal)
	fmt.Println(value)

	IntValue, _ := strconv.Atoi("20000") //cara mudah mengkonversi string ke int
	fmt.Println(IntValue)
	return ""
}
