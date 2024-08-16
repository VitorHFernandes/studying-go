package main

import "fmt"

func main() {
	var var1 string = "Variable 1" //TODO -> Explicit
	var2 := "Variable 2"           //TODO -> Implicit Infering types

	fmt.Println(var1)
	fmt.Println(var2)

	var (
		var3 string = "Variable 3"
		var4 string = "Variable 4"
	)

	fmt.Println(var3, var4)

	var5, var6 := "Variable 5", "Variable 6"

	fmt.Println(var5, var6)

	const constant1 string = "Const 1"

	fmt.Println(constant1)

	var5, var6 = var6, var5
	fmt.Println(var5, var6)
}
