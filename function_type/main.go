package function_type

import "fmt"

type Filter func(string) string //membuat alias untuk function

func aSelole(firstName string, lastName string, midName string, bornTo string) {

	fmt.Println("Alooo", firstName, lastName, midName, bornTo)

}

func functionAsValue(name string) string {
	return "Function as value " + name
}

func returnValue(name string) string {
	// result := "Hello"
	// return result
	if name == "" {
		return "Hello Cuy"
	} else {
		return "Hello " + name
	} // return dalam if statement
}

func returnMultiValue() (string, string) {
	return "Gibs", "Utee"
}

// Penggunaan return yang tidak perlu disebutkan value nya, karena sudah ada diatas atau dijadikan parameter
func returnNameValue() (firstname string, middlename string, lastname string) {
	firstname = "Gibs"
	middlename = "Utee"
	lastname = "Gibut"
	return
}

func variadicFunction(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func funcAsParam(names string, filter Filter) {
	nameFiltered := filter(names)
	fmt.Println("Hello", nameFiltered)
}

func funcAsParam1(name string) string {
	if name == "Anjing" {
		return "..."
	} else if name == "Goblok" {
		return "..."
	} else {
		return name
	}
}

func main() {
	for i := 0; i < 10; i++ {
		firstName := "Meylda"
		lastName := "Utee"
		aSelole(firstName, lastName, "Tralala", "Depok")
		aSelole("Gibran", "Mahardika", "Gibs", "Bekasi")
	}

	// Proses pengambilang value ada 2 seperti dibawah
	// yg 1 cara panjang
	result := returnValue("Gibs")
	fmt.Println(result)
	// yg 2 cara pendek
	fmt.Println(returnValue(""))

	// Proses return multiple
	beginName, endName := returnMultiValue()
	fmt.Println(beginName)
	fmt.Println(endName)

	// proses return value yang dimana hanya menggunakan satu parameter
	startName, _ := returnMultiValue()
	fmt.Println(startName)

	// proses return value menggunakan name yg berbeda, name atau parameter
	a, b, c := returnNameValue()
	fmt.Println("return dengan variable a = ", a)
	fmt.Println("return dengan variable b = ", b)
	fmt.Println("return dengan variable c = ", c)

	// proses variadic function yang ditandai dengan titik tiga (...)
	total := variadicFunction(3, 4, 1, 2, 3, 6, 7, 5, 8, 67, 9, 0)
	fmt.Println("Hasil Variadic Function pertama = ", total)

	// Jika mempunyai slice dapat di overwrite nilainya menjadi seperti dibawah ini
	sliceFunc := []int{1123, 123, 123435, 45567, 243135, 456345}
	total = variadicFunction(sliceFunc...)
	fmt.Println(total)

	// function yang dijadikan value dalam variable
	funcValue := functionAsValue

	funcResult := funcValue("hola madefaka")
	fmt.Println("ini yang dijadikan value = ", funcResult)
	fmt.Println(functionAsValue("ini langsung tanpa dijadikan value = hola madefaka"))

	// function as parameter
	funcAsParam("Gibs", funcAsParam1)   //function funcAsParam untuk membuat function parameter filter
	funcAsParam("Anjing", funcAsParam1) //sedangkan funcAsParam1 untuk membuat kondisional filternya
	funcAsParam("Goblok", funcAsParam1)
}
