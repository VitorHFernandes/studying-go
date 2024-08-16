package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func nomes(nomes ...string) string {
	total2 := ""
	for _, primeiroNome := range nomes {
		total2 += primeiroNome
	}

	return total2
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalSoma := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	totalNome := nomes("Vítor", "Marcelo", "Izadora")

	escrever("Olá mundo!", 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)

	fmt.Println(totalSoma)
	fmt.Println(totalNome)
}
