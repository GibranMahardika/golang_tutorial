package closure

import "fmt"

func ClosureMain() string {
	// Closure adalah kemampuan sebuah function untuk berinteraksi dengan data-data disekitarnya dalam scope yg sama, scope dalam artian file
	// jika salah mengimplementasikan closure, makan akan bisa mengubah data yang seharusnya tidak diubah
	// kalo di python itu disebutnya global variable, yang dimana variable tersebut bisa diakses dalam function2 lainnya

	var title string = "closure - closure"
	fmt.Println(title)
	// variable ini disebut scope global meski dalam function ClosureMain
	name := "Gibs"
	counter := 0

	increment := func() {
		// variable dibawah ini disebut scope function increment
		name := "aselole"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
	return title
}
