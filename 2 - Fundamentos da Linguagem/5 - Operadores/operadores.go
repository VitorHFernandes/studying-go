package main

import "fmt"

func main() {
	// ARITIMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	//	var num1 int16 = 10
	//	var num2 int32 = 25

	// Nao posso trabalhar com dados de tipos diferentes!
	//soma2 := num1 + num2
	//fmt.Println(soma2)

	// FIM DOS ARITIMETICOS

	// ATRIBUIÇÃO
	var var1 string = "String"
	var2 := "String2"

	fmt.Println(var1, var2)
	// FIM DOS OPERADORES DE ATRIBUIÇÕES

	// OPERADORES RELACIONAIS
	fmt.Println(3 > 2)
	fmt.Println(3 >= 2)
	fmt.Println(3 == 2)
	fmt.Println(3 <= 2)
	fmt.Println(3 > 2)
	fmt.Println(3 < 2)
	fmt.Println(3 != 2)
	// FIM DOS RELACIONAIS

	//OPERADORES LOGICOS
	fmt.Println(true && false) // operador E
	fmt.Println(true || false) // Operador OU
	fmt.Println(!false)        // negação
	// FIM OPERADORES LÓGICOS

	// OPERADORES UNÁRIOS
	numero := 10
	numero++
	numero += 15 // numero = numero + 15
	numero--
	numero -= 20 // numero - 20
	numero *= 3  // numero * 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)

	// FIM OPERADORES UNÁRIOS

	// OPERADORES TERNARIOS
	// Não existe, portanto, será necessário usar if/else

	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)

}
