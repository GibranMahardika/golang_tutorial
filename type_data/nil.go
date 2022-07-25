package type_data

import "fmt"

// func NilImplementation() map[string]string {
// 	var person map[string]string = nil
// 	fmt.Println(person)
// 	return person
// }

func NewNilImplementation(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func NewOtherNilImplementation() map[string]string {
	var person map[string]string = NewNilImplementation("")

	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Print(person)
	}
	return person
}
