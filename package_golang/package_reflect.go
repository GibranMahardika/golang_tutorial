package package_golang

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"` //ini Tag
}

type Sample1 struct {
	Name        string //tanpa menggunakan Tag
	Description string
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func PackageReflect() string {
	sample := Sample{"Gibran"}

	var sampleType reflect.Type = reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println()

	fmt.Println("Penggunaan Tag pada Name string")
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))
	fmt.Println(sampleType.Field(0).Tag.Get("min"))
	fmt.Println()

	fmt.Println("Penggunaan Tag dengan validasi")
	sample.Name = ""
	fmt.Println(IsValid(sample))

	sample1 := Sample1{"", ""}
	fmt.Println(IsValid(sample1)) //hasilnya akan true, karena tidak menggunakan Tag
	return ""
}
