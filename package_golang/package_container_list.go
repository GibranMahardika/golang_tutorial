package package_golang

import (
	"container/list"
	"fmt"
)

func PackageContainerList() string {
	data := list.New()

	data.PushBack("Gibran")
	data.PushBack("Mahardika")
	data.PushFront("Gibs")

	// fmt.Println(data.Front().Value)
	// fmt.Println(data.Back().Value)

	// dari depan ke belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// dari belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}

	return ""
}
