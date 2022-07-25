package package_golang

import (
	"fmt"
	"time"
)

func PackageTime() string {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Day())
	fmt.Println(now.Month())
	fmt.Println(now.Year())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(1997, 5, 24, 12, 12, 12, 12, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02" //"2006-01-02" format default golang // time.RFC3339
	fmt.Println(layout)
	parse, _ := time.Parse(layout, "1996-01-02")
	fmt.Print(parse)

	return ""
}
