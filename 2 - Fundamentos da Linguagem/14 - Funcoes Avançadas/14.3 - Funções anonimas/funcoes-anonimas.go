package main

import "fmt"

func main() {
	func(texto string) {
		fmt.Println("Olá", texto)
	}("Vítor")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Vítor")

	fmt.Println(retorno)
}
