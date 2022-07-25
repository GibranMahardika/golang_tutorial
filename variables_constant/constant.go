package variables_constant

import "fmt"

func ConstantStyle() string {
	// constant merupakan sebuah variable yang dimana value nya tidak dapat diubah2
	// berbeda dengan variable yang ditandai dengan var, deklarasi constant menggunakan const
	const asoy string = "Deklarasi constant asoy dengan penulisan type data nya "
	const asoy1 = "Deklarasi constant asoy1 tanpa penulisan type data nya "
	const value = 10
	const (
		asoy2 = "Deklarasi multiple constant pada variable asoy2"
		asoy3 = "Deklarasi multiple constant pada variable asoy3"
	)
	fmt.Println(asoy)
	fmt.Println(asoy1)
	fmt.Println(asoy2)
	fmt.Println(asoy3)
	fmt.Println("Deklarasi constant age dengan type data int yang tanpa ada penulisan type data nya = ", value)
	return asoy
}
