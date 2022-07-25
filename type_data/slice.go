package type_data

import "fmt"

func Slice() string {

	// TUTORIAL SLICE
	// NOTE : JIKA SLICE DIUBAH DATANYA, AKAN MEMPENGARUHI ARRAY, KARENA SLICE MENGAMBIL DATA DARI ARRAY,
	// BEGITUPUN SEBALIKNYA, JIKA ARRAY DIUBAH DATANYA, MAKA SLICE JUGA AKAN BERUBAH HASIL DARI PRINT NYA.

	var title string = "type_data - slice"
	fmt.Println(title)

	var intSlice = [10]int{
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		0,
	}

	fmt.Println("Data array intSlice sebelum diubah dengan slice = ", intSlice)
	var slice1 = intSlice[4:7] //variable ini dibuat untuk mencari index pada array intSlice, yaitu dengan index ke 4 sampai dengan ke 7

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// CONTOH SLICE ATAU ARRAY YANG DIUBAH
	intSlice[5] = 6000 //Syntax ini untuk mengganti value dari array intSlice, yang dimana index 5 itu diganti valuenya menjadi 6000
	fmt.Println("Data array intSlice index 5 yang diubah = ", slice1)

	slice1[0] = 5000
	fmt.Println("Data intSlice dengan index 0 diubah = ", slice1)
	//NOTE: variable slice1 dengan index 0 itu maksudnya,
	// slice1 itu ngambil data dari intSlice dengan index 4 sampai 7,
	// jadi slice1 index 0 itu sama dengan intSlice index 4

	fmt.Println("Data array yang sudah diubah semuanya = ", intSlice)

	// CONTOH SLICE MENGGUNAKAN APPEND
	// Jika slice append melewati kapasitas dari array, maka slice tersebut akan membuat array baru,
	// yang tidak ada sangkut pautnya dengan array yg lama
	var slice2 = intSlice[9:]
	fmt.Println("Data intSlice pada variable slice2 = ", slice2)

	var slice3 = append(slice2, 11000)
	fmt.Println("Contoh Append dengan menggunakan variable slice baru, yaitu slice3 = ", slice3)

	// Setelah di append, karena membuat array baru, maka sebagai bukti harus di print
	fmt.Println("Slice2 = ", slice2)
	fmt.Println("intSlice = ", intSlice)
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

	return title
}
