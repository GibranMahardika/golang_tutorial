package package_golang

import (
	"fmt"
	"math"
)

func PackageMath() string {
	fmt.Println(math.Round(1.7))
	fmt.Println(math.Round(1.3))
	fmt.Println(math.Round(1.5))

	fmt.Println(math.Floor(1.7))
	fmt.Println(math.Ceil(1.7))

	fmt.Println(math.Max(10, 20))
	fmt.Println(math.Min(10, 20))
	return ""
}
