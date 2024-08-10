package main

import "fmt"

func main() {
	fmt.Println("If/else")

	num := 10

	if num > 15 {
		fmt.Println("Maior que 15")
	} else if num < -10 {
		fmt.Println("Menor que -10")
	} else {
		fmt.Println("Entre 0 e 10")
	}

	if outroNumero := num; outroNumero > 0 {
		fmt.Println("Numero é maior que zero!")
	}

	//fmt.Println(outroNumero) => Não definida, está limitada ao escopo do if/else
}
