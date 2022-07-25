package main

import (
	"fmt"
	"udemy_golang_course/package_golang"
)

func main() {

	// a := add.Add(1, 2)
	// fmt.Println(a)

	// fmt.Println()
	// type_number := type_data.TypeNumber()
	// fmt.Println(type_number)

	// fmt.Println()
	// type_bool := type_data.TypeBoolean()
	// fmt.Println(type_bool)

	// fmt.Println()
	// type_string := type_data.TypeString()
	// fmt.Println(type_string)

	// fmt.Println()
	// variable_style := variables_constant.VariablesStyle()
	// fmt.Println(variable_style)

	// fmt.Println()
	// constant_style := variables_constant.ConstantStyle()
	// fmt.Println(constant_style)

	// fmt.Println()
	// conversion_Tdata := type_data.ConversionTypeData()
	// fmt.Println(conversion_Tdata)

	// fmt.Println()
	// type_declaration := type_data.TypeDeclaration()
	// fmt.Println(type_declaration)

	// fmt.Println()
	// math_ops := math_operation.Math()
	// fmt.Println(math_ops)

	// fmt.Println()
	// comp_ops := math_operation.CompOperation()
	// fmt.Println(comp_ops)

	// fmt.Println()
	// bool_ops := operasi_operator.BooleanOperator()
	// fmt.Println(bool_ops)

	// fmt.Println()
	// array := type_data.Array()
	// fmt.Println(array)

	// fmt.Println()
	// slice := type_data.Slice()
	// fmt.Println(slice)

	// fmt.Println()
	// maps := type_data.Map()
	// fmt.Println(maps)

	// fmt.Println()
	// if_statement := if_else_switch.IfExpression()
	// fmt.Println(if_statement)

	// fmt.Println()
	// switch_statement := if_else_switch.SwitchExpression()
	// fmt.Println(switch_statement)

	// fmt.Println()
	// for_loop := for_loop.ForLoop()
	// fmt.Println(for_loop)

	// fmt.Println()
	// function_first := function_type.FunctionFirst()
	// fmt.Println(function_first)

	// fmt.Println()
	// func_parameter := function_type.FunctionParameter("Gibran ", "Mahardika")
	// fmt.Println(func_parameter)

	// fmt.Println()
	// func_return_value := function_type.FunctionReturnValue("")
	// fmt.Println(func_return_value)
	// fmt.Println(function_type.FunctionReturnValue("Gibs"))

	// fmt.Println()
	// func_multiple_return_value, func_multiple_return_value1, func_multiple_return_value2 := function_type.FunctionMultipleReturnValue()
	// _, _, func_multiple_return_value3 := function_type.FunctionMultipleReturnValue()
	// //jika tidak membutuhkan return, hanya menggunakan underscore (_) contoh diatas
	// fmt.Println("func_multiple_return_value = ", func_multiple_return_value)
	// fmt.Println("func_multiple_return_value1 = ", func_multiple_return_value1)
	// fmt.Println("func_multiple_return_value2 = ", func_multiple_return_value2)
	// fmt.Println("func_multiple_return_value3 = ", func_multiple_return_value3)
	// fmt.Println("func_multiple_return_value1 = ", func_multiple_return_value, ", func_multiple_return_value1 = ", func_multiple_return_value1, ", func_multiple_return_value2 = ", func_multiple_return_value2)

	// fmt.Println()
	// funcTitle, funcFirstName, funcLastName, funcSlackName := function_type.FunctionReturnNamedValue()
	// fmt.Println("Title = ", funcTitle)
	// fmt.Println("First Name = ", funcFirstName)
	// fmt.Println("Last Name = ", funcLastName)
	// fmt.Println("Slack Name = ", funcSlackName)
	// fmt.Println("Title = ", funcTitle,
	// 	"First Name = ", funcFirstName,
	// 	"Last Name = ", funcLastName,
	// 	"Slack Name = ", funcSlackName)

	// fmt.Println()
	// func_variadic := function_type.FunctionVariadic(1, 2, 3, 4, 5)
	// fmt.Println("Variadic Function secara manual = ", func_variadic)

	// func_vard_slice := []int{1, 3, 2, 44, 55}
	// func_variadic = function_type.FunctionVariadic(func_vard_slice...)
	// fmt.Println("Variadic Function menggunakan slice = ", func_variadic) //Implementasi variadic function ketika mempunyai slice baru atau lama

	// fmt.Println()
	// func_as_val := function_type.FunctionAsValue
	// func_as_val1 := func_as_val("1")
	// fmt.Println(func_as_val1)

	// fmt.Println()
	// func_as_param_filter := function_type.FunctionAsParameter1
	// func_as_param := function_type.FunctionAsParameter("Gibs", func_as_param_filter)
	// fmt.Println(func_as_param)

	// fmt.Println()
	// // Cara ke 1
	// blacklist := func(name string) bool {
	// 	return name == "admin"
	// }
	// function_type.FunctionAnonymous("admin", blacklist)
	// function_type.FunctionAnonymous("Gibs", blacklist)
	// // Cara ke 2
	// function_type.FunctionAnonymous("root", func(name string) bool {
	// 	return name == "root"
	// })
	// function_type.FunctionAnonymous("Gibs", func(name string) bool {
	// 	return name == "root"
	// })

	// fmt.Println()
	// fmt.Println("function_type - recursive function")
	// factorial_loop := function_type.FactorialLoop(5)
	// fmt.Println(factorial_loop)
	// fmt.Println(5 * 4 * 3 * 2 * 1)

	// factorial_recursive := function_type.FactorialRecursive(10)
	// fmt.Println(factorial_recursive)

	// fmt.Println()
	// closure_imp := closure.ClosureMain()
	// fmt.Println(closure_imp)

	// fmt.Println()
	// defer_func1 := function_type.RunApplication(100)
	// // jika dibagi nol 0 akan eror
	// fmt.Println(defer_func1)

	// fmt.Println()
	// panic_func := function_type.RunApp(true)
	// fmt.Println(panic_func)

	// fmt.Println()
	// type_struct := type_data.StructImplementation()
	// fmt.Println(type_struct)

	// fmt.Println()
	// fmt.Println("type_data - interface")
	// var gibs type_data.Person
	// gibs.Name = "Gibs"
	// gibs.Number = 12345
	// type_data.SayHello(gibs)

	// fmt.Println()
	// fmt.Println("type_data - interface_empty")
	// var data interface{} = type_data.Ups(1)
	// fmt.Println(data)

	// fmt.Println()
	// fmt.Println("type_data - nil")
	// // type_nil := type_data.NilImplementation()
	// // fmt.Println(type_nil)

	// type_new_nil := type_data.NewOtherNilImplementation()
	// fmt.Println(type_new_nil)

	// fmt.Println()
	// fmt.Println("type_data - interface_error")
	// hasil, err := type_data.Pembagi(100, 2)
	// if err == nil {
	// 	fmt.Println("Hasil", hasil)
	// } else {
	// 	fmt.Println("Error", err.Error())
	// }

	// fmt.Println()
	// fmt.Println("type_data - type_assertions")
	// // var typeAssertion = type_data.TypeAssertionsImplementation()
	// var typeAssertionSwitch = type_data.TypeAssertionsSwitch()
	// // fmt.Println(typeAssertion)
	// fmt.Println(typeAssertionSwitch)

	// //====================================================

	// fmt.Println()
	// fmt.Println("pointer - pointer")
	// passingValue := pointer.PassingValue()
	// passingReference := pointer.PassingReference()
	// fmt.Println(passingValue)
	// fmt.Println()
	// fmt.Println(passingReference)

	// fmt.Println()
	// fmt.Println("pointer - pointer_function")
	// var alamat = pointer.Address1{
	// 	City:     "Bekasi",
	// 	Province: "Jawa Barat",
	// 	Country:  "",
	// }
	// pointer.ChangeCountryIndonesia(&alamat)
	// fmt.Println(alamat)

	// fmt.Println()
	// pointer_method := pointer.PointerMethod()
	// fmt.Println(pointer_method)

	// var gibss = pointer.Man{"Gibran"}
	// gibss.Married()
	// fmt.Println(gibss.NameMan)

	// var utee = pointer.Woman{"Utee"}
	// utee.Married()
	// fmt.Println(utee.NameWoman)
	// fmt.Println()

	// =============================================

	fmt.Println("package_golang - package_initialization")
	packageInitialization := package_golang.PackageInitialization()
	fmt.Print(packageInitialization)

	fmt.Println("package_golang - package_os")
	var packageOs = package_golang.PackageOs()
	fmt.Println(packageOs)

	fmt.Println("package_golang - package_flag")
	packageFlag := package_golang.PackageFlag()
	fmt.Println(packageFlag)

	fmt.Println("package_golang - package_flag")
	packageStrings := package_golang.PackageStrings()
	fmt.Println(packageStrings)

	fmt.Println("package_golang - package_strconv")
	packageStrconv := package_golang.PackageStrconv()
	fmt.Println(packageStrconv)

	fmt.Println("package_golang - package_math")
	packageMath := package_golang.PackageMath()
	fmt.Println(packageMath)

	fmt.Println("package_golang - package_math")
	packageContainerList := package_golang.PackageContainerList()
	fmt.Println(packageContainerList)

	fmt.Println("package_golang - package_container_ring")
	packageContainerRing := package_golang.PackageContainerRing()
	fmt.Println(packageContainerRing)

	fmt.Println("package_golang - package_sort")
	packageSort := package_golang.PackageSort()
	fmt.Println(packageSort)
}
