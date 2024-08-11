package main

import (
	"fmt"
	"time"
)

func escrever(texto string, channel chan string) {
	//	time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		channel <- texto // Envio do valor para o canal
		time.Sleep(time.Second)
	}

	close(channel) // Fecha o canal ao sair do for. //! EVITA DEADLOCK
}

func main() {
	channel := make(chan string)

	go escrever("Olá mundo!", channel)

	fmt.Println("Depois da função escrever começar a ser executada")

	/*
		==> EXEMPLO #01 <==
		for {
			mensagem, aberto := <-channel // Espero o valor ser inserido no canal
			if !aberto {                  //TODO => Se o canal não estiver aberto, saí do LOOP
				break
			}
			fmt.Println(mensagem)
		}

	*/

	for mensagem := range channel {
		fmt.Println(mensagem)
	}
	fmt.Println("Fim do programa!")
}
