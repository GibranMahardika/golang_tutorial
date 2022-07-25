package type_data

import "errors"

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}
