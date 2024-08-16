package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estrutura de repetição")
	/*
		i := 0

		for i < 10 {
			i++
			fmt.Println("Incrementando i")
			time.Sleep(time.Second)
		}

		for j := 0; j < 10; j++ {
			fmt.Println("Incrementando j", j)
			time.Sleep(time.Second)
		}
	*/
	nomes := [3]string{"João", "Davi", "Lucas"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}
	fmt.Println("------")
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}
	fmt.Println("------")
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra)
	}

	fmt.Println("------")
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
