package function_type

func FactorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= 1
	}
	return result
}

func FactorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * FactorialRecursive(value-1)
	}
}
