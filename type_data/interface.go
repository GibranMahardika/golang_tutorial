package type_data

import "fmt"

type HasName interface {
	GetName() string
	GetNumber() int
}

func SayHello(people HasName) {
	fmt.Println("Hello ", people.GetName())
	fmt.Println("Can I get your Number ?  ", people.GetNumber())
}

type Person struct {
	Name   string
	Number int
}

func (person Person) GetName() string {
	return person.Name
}

func (personNumber Person) GetNumber() int {
	return personNumber.Number
}
