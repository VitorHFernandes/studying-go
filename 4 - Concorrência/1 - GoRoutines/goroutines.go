package main

import (
	"fmt"
	"time"
)

// CONCORRENCIA != PARALELISMO

func main() {
	go escrever("Olá mundo!") // goroutine, inicie e passa pra sentença abaixo
	escrever("Programando em GO!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
