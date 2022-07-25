package variables_constant

import "fmt"

func VariablesStyle() string {
	// cara-cara penulisan variable
	var names string //tanpa ada pendeklarasian
	names = "Deklarasi variable names secara manual"
	var names1 string = "Deklarasi variable names1 secara langsung"
	names2 := "Deklarasi sebuah variable names2 secara langsung yang tidak perlu menulis var dan menulis type data nya, tetapi hanya satu kali saja karena menggunakan :="
	var names3 = "Deklarasi variable names3 tapi tanpa harus menulis type data nya"
	var age = 1

	var (
		names4 = "Deklarasi variable names4 yang disatukan dalam bentuk string"
		names5 = "Deklarasi variable names5 yang disatukan dalam bentuk string"
	) // cara deklarasi variable dijadikan satu

	fmt.Println(names)
	fmt.Println(names1)
	fmt.Println(names2)
	fmt.Println()
	fmt.Println(names3)
	fmt.Println(names4)
	fmt.Println(names5)
	fmt.Println("ini adalah variable age, dengan type data int = ", age)

	return names
}
