package package_golang

var connection string

func init() {
	connection = "MySQL"
}

func PackageInitialization() string {
	return connection
}
