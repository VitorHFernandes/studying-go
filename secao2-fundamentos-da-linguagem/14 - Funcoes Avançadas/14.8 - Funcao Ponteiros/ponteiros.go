package main

import "fmt"

func inverterSinal(n1 int) int {
	return n1 * -1
}

func inverterSinalPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)

	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 40

	inverterSinalPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
