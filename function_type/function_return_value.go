package function_type

func FunctionReturnValue(name string) string {
	if name == "" {
		return "Nama Kosong"
	} else {
		return "Hello " + name
	}
}
