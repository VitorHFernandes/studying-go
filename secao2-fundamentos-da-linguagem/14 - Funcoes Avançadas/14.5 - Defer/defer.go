package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função #01")
}

func funcao2() {
	fmt.Println("Executando a função #02")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. O resultado será retornado abaixo:") // Adiar até o último momento possível
	fmt.Println("Entrando na função para verificar se o aluno está aprovado!")

	media := (n1 + n2) / 2

	return media >= 6
}

func main() {
	funcao1()
	funcao2()

	defer fmt.Println(alunoAprovado(1, 10))
}
