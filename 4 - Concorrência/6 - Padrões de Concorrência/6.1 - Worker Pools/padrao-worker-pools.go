package main

import "fmt"

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	// Quantos forem necessários, depende da quantidade de núcleos do processador.
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for j := 0; j < 45; j++ {
		resultado := <-resultados
		fmt.Println(resultado)

	}
}

// <-chan: indica que o canal só recebe informações
// chan<-: indica que o canal só envia informações
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
