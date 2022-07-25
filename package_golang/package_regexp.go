package package_golang

import (
	"fmt"
	"regexp"
)

func PackageRegexp() string {
	var regex *regexp.Regexp = regexp.MustCompile("g([a-z])b")

	fmt.Println(regex.MatchString("gib"))
	fmt.Println(regex.MatchString("gab"))
	fmt.Println(regex.MatchString("guab"))

	cari := regex.FindAllString("gib gab g3b g0b gbi", -1)
	fmt.Println(cari)
	return ""
}
