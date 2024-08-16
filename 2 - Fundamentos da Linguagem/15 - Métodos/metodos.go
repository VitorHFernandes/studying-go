package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// "u" seria como um alias, um apelido
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func (u usuario) maiorDeIdade() {
	if u.idade < 18 {
		fmt.Printf("O usuário %s é menor de idade! Ele tem apenas %d anos.\n", u.nome, u.idade)
	} else {
		fmt.Printf("O usuário %s é maior de idade! Ele já tem %d anos.\n", u.nome, u.idade)
	}
}

func main() {
	fmt.Println("MÉTODO")
	usuario1 := usuario{"Vítor", 21}

	usuario2 := usuario{"Izadora", 13}

	usuario1.salvar()
	usuario2.salvar()

	usuario1.maiorDeIdade()
	usuario2.maiorDeIdade()

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
