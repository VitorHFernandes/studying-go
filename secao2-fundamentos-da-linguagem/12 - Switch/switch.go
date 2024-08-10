package main

import "fmt"

func diaSemana(num uint8) string {
	switch num {
	case 1:
		return "Domingo"

	case 2:
		return "Segunda-Feira"

	case 3:
		return "Terça-Feira"

	case 4:
		return "Quarta-Feira"

	case 5:
		return "Quinta-Feira"

	case 6:
		return "Sexta-Feira"

	case 7:
		return "Sábado"
	default:
		return "Número inválido!"
	}
}

func diaSemana2(num int) string {
	var diaSemaninha string
	switch {
	case num == 1:
		diaSemaninha := "Domingo"
		fallthrough // Joga para o próximo resultado
	case num == 2:
		diaSemaninha := "Segunda-Feira"
	default:
		diaSemaninha := "Numero invalido!"
	}
	return diaSemaninha
}

func main() {
	fmt.Println(diaSemana(1))
	dia2 := diaSemana2(1)
	fmt.Println(dia2)
}
