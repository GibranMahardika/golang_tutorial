package operasi_operator

import "fmt"

func BooleanOperator() string {

	var title string = "operasi_operator - boolean_operator"
	fmt.Println(title)

	var ujian = 80
	var absensi = 60

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 80
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusAbsensi && lulusUjian

	fmt.Println(lulus)
	fmt.Println(ujian >= 80 && absensi >= 80)
	return title
}
