package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	// Sequencia de Fibonacci
	// 1 1 2 3 5 8 13...
	posicao := 45
	for i := 0; i <= posicao; i++ {
		fmt.Println(fibonacci(uint(i)))
	}
}
