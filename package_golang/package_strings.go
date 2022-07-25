package package_golang

import (
	"fmt"
	"strings"
)

func PackageStrings() string {
	fmt.Println(strings.Contains("Gibran Mahardika", "Gibran"))
	fmt.Println(strings.Contains("Gibran Mahardika", "Dika")) //hasilnya akan false, karena contains menggunakan case sensitive

	fmt.Println(strings.Split("Gibran Mahardika", " "))

	fmt.Println(strings.ToLower("Gibran Mahardika"))
	fmt.Println(strings.ToUpper("Gibran Mahardika"))
	fmt.Println(strings.ToTitle("gibran mahardika"))

	fmt.Println(strings.Trim("     Gibran Mahardika   ", " "))
	fmt.Println(strings.ReplaceAll("Gibran Gibran Gibran", "Gibran", "Gibs"))
	return ""
}
