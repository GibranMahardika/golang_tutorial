package package_golang

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}
type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func PackageSort() string {
	users := []User{
		{"Gibran", 123},
		{"Gibs", 4323},
		{"Mahardika", 67},
		{"Dika", 354},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)
	return ""
}
