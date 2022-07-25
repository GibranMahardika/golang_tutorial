package pointer

type Man struct {
	NameMan string
}

type Woman struct {
	NameWoman string
}

func (man *Man) Married() {
	man.NameMan = "Mr. " + man.NameMan
}

func (woman *Woman) Married() {
	woman.NameWoman = "Mrs. " + woman.NameWoman
}

func PointerMethod() string {
	var title = "pointer - pointer_method"
	return title
}
