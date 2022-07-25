package package_golang

import (
	"flag"
	"fmt"
)

func PackageFlag() string {
	var host *string = flag.String("host", "localhost", "Put your Database host")
	var user *string = flag.String("user", "localhost", "Put your Database user")
	var pass *string = flag.String("password", "localhost", "Put your Database password")
	//var menggunakan pointer pada type data nya

	flag.Parse()
	fmt.Println("Host : ", *host)
	fmt.Println("Username : ", *user)
	fmt.Println("Password : ", *pass)
	//pada saat print harus menggunakan pointer agar data nya keambil
	return ""
	//cara run di cmd .. go run main.go -host=localhost -user=root -password=root
}
