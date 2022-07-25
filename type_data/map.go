package type_data

import "fmt"

func Map() string {
	// Type data MAP

	var title string = "type_data - map"
	fmt.Println(title)

	var person map[string]string = map[string]string{
		"name":    "Gibran",
		"address": "Bekasi",
	} //deklarasi panjang dalam pembuatan map
	fmt.Println(person)

	person1 := map[string]string{
		"name":   "Gibran",
		"pacar1": "Meylda",
	} //deklarasi singkat dalam pembuatan map

	fmt.Println(person1)
	fmt.Println(person1["name"])
	fmt.Println(person1["pacar1"])
	fmt.Println(person1["pacar2"])

	person1["pacar2"] = "Utee" //syntax nambah data pada map

	var book map[string]string = make(map[string]string) //syntax membuat map
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
