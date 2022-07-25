package type_data

import "fmt"

func ConversionTypeData() string {

	var title string = "type_data - conversion"
	fmt.Println(title)

	var val32 int32 = 1000000      // Deklarasi variable val32 dengan value 1jt
	var val64 int64 = int64(val32) // Deklarasi variable val64 untuk menyimpan hasil conversi variable val32 yang dimana val32 itu hanya 32byte menjadi 64byte
	var val8 int8 = int8(val32)    // Deklarasi variable val8 untuk menyimpan hasil conversi variable val32 yang dimana val32 itu hanya 32byte menjadi 8byte

	fmt.Println(val32)
	fmt.Println(val64)
	fmt.Println(val8)

	// cara conversi byte menjadi string
	var nameString = "Gibs"
	var e byte = nameString[0] //e merupakan variable / alias dari tipe data byte, yang dimana mendeklarasikan dirinya untuk mencari index 0 dari value variable nameString
	var eString = string(e)    //sedangkan eString merupakan variable untuk mengkonversi variable e yang sebelumnya byte menjadi string

	fmt.Println(nameString)
	fmt.Println(e)
	fmt.Println(eString)
	return title
}
