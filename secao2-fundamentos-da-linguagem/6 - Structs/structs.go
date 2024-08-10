package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint16
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Vitor"
	u.idade = 21
	fmt.Println(u)

	enderecoEx := endereco{"R. Tuiuti", 3222}

	usuario2 := usuario{"Davi", 21, enderecoEx}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "David Guetta"}
	fmt.Println(usuario3)
}
