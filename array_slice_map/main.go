package main

import "fmt"

func main() {
	// TUTORIAL ARRAY
	// apakah array bisa melakukan increment?

	fmt.Println("TUTORIAL ARRAY")
	var stringArr [3]string //ini deklarasi array secara manual
	var intArr [10]int
	var IntArr = [10]int{
		10, 20, 30, 40, 50, 60, 70, 80, 90, 100,
	} // ini contoh deklarasi secara langsung

	stringArr[0] = "Gibs"
	stringArr[1] = "Gibran"
	stringArr[2] = "Mahardika"

	intArr[0] = 0
	intArr[1] = 1
	intArr[2] = 2
	intArr[3] = 3
	intArr[4] = 4
	intArr[5] = 5
	intArr[6] = 6
	intArr[7] = 7
	intArr[8] = 8
	intArr[9] = 9

	fmt.Println(stringArr)
	fmt.Println(stringArr[0])
	fmt.Println(stringArr[1])
	fmt.Println(stringArr[2])
	fmt.Println(intArr)
	fmt.Println(intArr[0])
	fmt.Println(intArr[1])
	fmt.Println(intArr[2])
	fmt.Println(intArr[3])
	fmt.Println(intArr[4])
	fmt.Println(intArr[5])
	fmt.Println(intArr[6])
	fmt.Println(intArr[7])
	fmt.Println(intArr[8])
	fmt.Println(intArr[9])

	fmt.Println("Jumlah stringArr = ", len(stringArr))
	fmt.Println("Jumlah intArr = ", len(intArr))
	fmt.Println("Jumlah IntArr = ", len(IntArr))

	fmt.Println("Contoh array langsung = ", IntArr)

	// TUTORIAL SLICE
	// NOTE : JIKA SLICE DIUBAH DATANYA, AKAN MEMPENGARUHI ARRAY, KARENA SLICE MENGAMBIL DATA DARI ARRAY,
	// BEGITUPUN SEBALIKNYA, JIKA ARRAY DIUBAH DATANYA, MAKA SLICE JUGA AKAN BERUBAH HASIL DARI PRINT NYA.

	fmt.Println("Data array IntArr sebelum diubah dengan slice = ", IntArr)
	var slice1 = IntArr[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	// CONTOH SLICE ATAU ARRAY YANG DIUBAH
	IntArr[5] = 6000
	fmt.Println("Data array IntArr yang diubah = ", slice1)

	slice1[0] = 5000
	fmt.Println("Data array IntArr yang diubah dengan slice = ", slice1)

	fmt.Println("Data array yang sudah diubah semuanya = ", IntArr)

	// CONTOH SLICE MENGGUNAKAN APPEND
	// Jika slice append melewati kapasitas dari array, maka slice tersebut akan membuat array baru,
	// yang tidak ada sangkut pautnya dengan array yg lama
	var slice2 = IntArr[9:]
	fmt.Println("Data IntArr pada variable slice2 = ", slice2)

	var slice3 = append(slice2, 11000)
	fmt.Println("Contoh Append dengan menggunakan variable slice baru, yaitu slice3 = ", slice3)

	// Setelah di append, karena membuat array baru, maka sebagai bukti harus di print
	fmt.Println("Slice2 = ", slice2)
	fmt.Println("IntArr = ", IntArr)
	// Benar kan, liat aja di terminal

	// Tutorial Membuat Slice baru
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Gibran"
	newSlice[1] = "Mahardika"

	fmt.Println(newSlice)
	fmt.Println("Panjang dari array slice = ", len(newSlice))
	fmt.Println("Kapasitas dari array slice = ", cap(newSlice))

	// Tutorial Copy Slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println("This is copySlice = ", copySlice)

	// Percobaan pada slice1, tadi coba langsung dengan array nya malah eror, entah kenapa besok tanyain ajadh
	copyIntArraySlice := make([]int, len(slice1), cap(slice1))
	copy(copyIntArraySlice, slice1)
	fmt.Println("Copy from slice1 = ", copyIntArraySlice)

	// Cara membedakan Array dengan Slice
	iniArray := [5]int{1, 2, 3, 4, 5} //array harus ada length didalam [] braket
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println("iniArray = ", iniArray)
	fmt.Println("iniSlice = ", iniSlice)

	// Type data MAP
	var person map[string]string = map[string]string{
		"name":    "Gibran",
		"address": "Bekasi",
	}
	fmt.Println(person)

	person1 := map[string]string{
		"name":   "Gibran",
		"pacar1": "Meylda",
	} //type singkat dalam pembuatan map

	person1["pacar2"] = "Utee"

	fmt.Println(person1)
	fmt.Println(person1["name"])
	fmt.Println(person1["pacar1"])
	fmt.Println(person1["pacar2"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Sabar"
	book["author"] = "Gibran Mahardika"
	book["ciusan?"] = "Tapi boong"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ciusan?")

	fmt.Println(book)
	fmt.Println(len(book))

}
