package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("Teste")
	generica(1)
	generica(true)

	fmt.Println()
}
