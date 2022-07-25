package pointer

import "fmt"

type Address struct {
	City, Province, Country string
}

func PassingValue() string {

	var title1 string = "pointer - pointer - PassingValue"
	fmt.Println(title1)

	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println()
	return ""

}

func PassingReference() string {

	var title2 string = "pointer - pointer - PassingReference"
	fmt.Println(title2)

	address11 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address22 := &address11
	address33 := &address11
	var address44 *Address = new(Address) //Membuat record kosong pada type Address

	address22.City = "Depok"
	address44.City = "Unknown" //Menambahkan record pada key City dalam type Address

	// address22 = &Address{"Bekasi", "Planet Tata Surga", "Jauh Deh"} //Pointer yang menambah value dari Address yang hanya diakses oleh variable address22
	*address22 = Address{"Bekasi", "Planet Tata Surga", "Jauh Deh"} //Pointer yang mengubah value dari seluruh type Address

	fmt.Println(address11)
	fmt.Println(address22)
	fmt.Println(address33)
	fmt.Println(address44)
	return ""
}
