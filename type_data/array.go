package type_data

import "fmt"

func Array() string {
	// TUTORIAL ARRAY
	// apakah array bisa melakukan increment?
	var title string = "type_data - array"
	fmt.Println(title)

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
	return title
}
