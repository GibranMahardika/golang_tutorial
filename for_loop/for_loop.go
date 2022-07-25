package for_loop

import "fmt"

func ForLoop() string {

	var title string = "for_loop - for_loop"

	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	// contoh init post statement

	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println("Perulangan ke ", counter)
	// }

	slice := []string{"Gibs", "Utee"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	// For dengan Map
	person := make(map[string]string)
	person["name"] = "Gibs"
	person["job"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

	return title
}
