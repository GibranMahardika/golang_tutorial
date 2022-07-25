package main

import "fmt"

//break itu untuk menghentikan perulangan for
// continue itu untuk menghentikan perulangan saat ini dan lanjut ke perulangan berikutnya atau post statement

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan break ke ", i)
	}

	for z := 0; z < 10; z++ {
		if z%2 == 0 {
			continue
		}
		fmt.Println("Perulangan continue ke ", z)
	}
}
