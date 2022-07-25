package function_type

import (
	"fmt"
)

func endApp() {
	message := recover() //function recover ini digunakan untuk menangkap data panic
	if message != nil {
		fmt.Println("Error dengan message: ", message)
	}
	fmt.Println("Aplikasi Selesai")
}

func RunApp(error bool) bool {
	defer endApp()
	if error {
		panic("APLIKASI ERROR") // ini di eksekusi, langsung di catch oleh recover() yang ada di function endApp //kenapa bisa di catch? karena ada fungsi defer, itu untuk melempar value ketika panic //karena fungsi panic itu untuk mematikan suatu program
	} else {
		fmt.Println("Aplikasi berjalan")
	}
	return error
}
