package function_type

import "fmt"

// Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
func functionDefer() string {
	fmt.Println("Ini FunctionDefer")
	return "_"
}

func RunApplication(value int) int {
	defer functionDefer()
	// NOTES saat menggunakan defer disarankan diatas,
	// atau setelah penamaan function

	// defer digunakan ketika ada eror,
	// tapi function atas/lainnya mau di eksekusi, maka gunakan defer

	// cara testnya input nilai 0 pada file main.go, lalu run,
	// ketika ada eror function yang menggunakan fungsi defer akan tereksekusi

	fmt.Println("Run Application") // pertama eksekusi
	result := 10 / value
	fmt.Println("Result", result)
	// Logging()
	return int(value)
}
