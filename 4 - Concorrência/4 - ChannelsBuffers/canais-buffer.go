package main

import "fmt"

//** BUFFER É UM ESPAÇO NA MEMÓRIA FÍSICA DEDICADO A ARMAZENAR INFORMAÇÕES TEMPORARIAMENTE.
//TODO => Utilizar canais com buffer é uma alternativa quando se precisa utilizar canais dentro da mesma função, o que não é recomendado.

func main() {
	channel := make(chan string, 2)

	channel <- "Olá mundo! #01"
	channel <- "Olá mundo! #02"

	mensagem := <-channel
	mensagem2 := <-channel

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
