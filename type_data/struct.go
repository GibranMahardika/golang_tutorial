package type_data

import "fmt"

type Customer struct {
	Name, Address, Title string
	Age, PhoneNumber     int
}

type Babe struct {
	BabeName, BabeAddress string
	BabeAge, BabePhone    int
}

// Implementasi struct method

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello ", name, "My name is", customer.Name)
	fmt.Println("I live in ", customer.Address)
}

func (hot Babe) Flirting(name string) {
	fmt.Println("Hi im", name, "what is your name ? ", hot.BabeName)
	fmt.Println(hot.BabeName, ", That's a beautiful name ")
	fmt.Println("Where do you live ? ", hot.BabeAddress)
}

//

func StructImplementation() string {
	var Gibs Customer
	Gibs.Name = "Gibran Mahardika"
	Gibs.Address = "Bekasi"
	Gibs.Title = "type_data - struct"
	Gibs.Age = 26
	Gibs.PhoneNumber = 0007

	Gibs.sayHi("Panjul")
	fmt.Println()

	Utee := Babe{
		BabeName:    "Meylda Utee",
		BabeAddress: "PAL",
		BabeAge:     25,
		BabePhone:   0123456,
	}
	Utee.Flirting("Gibs")

	return ""

	// fmt.Println(Gibs)
	// fmt.Println(Gibs.Name)
	// fmt.Println(Gibs.Title)
	// fmt.Println(Gibs.Age)
	// fmt.Println(Gibs.PhoneNumber)

	// Meylda := Customer{
	// 	Name:        "Meylda Kurnia Emylia Putri",
	// 	Address:     "PAL",
	// 	Title:       "type_data - struct",
	// 	Age:         25,
	// 	PhoneNumber: 12345,
	// }
	// Meylda.sayHi("Panjul")

	// fmt.Println(Meylda)

	// var Putri Customer = Customer{"Putri Emylia Kurnia Meylda", "LAP", "type_data - struct", 52, 54321}
	// fmt.Println(Putri)

}
