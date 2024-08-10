package main

import (
	"errors"
	"fmt"
)

func main() {
	//	int, int8, int16, int32, int64
	// uint -> unsygned int (nao aceita numeros negativos)
	var num1 int = 127
	var num2 uint32 = 10000

	fmt.Println(num1)
	fmt.Println(num2)

	//alias
	//int32 => rune
	var num3 rune = 123456
	fmt.Println(num3)

	//int8 	=> byte
	var num4 byte = 10
	fmt.Println(num4)

	var numReal1 float32 = 123000000000000000000000000000000000000.35
	fmt.Println(numReal1)

	var numReal2 float64 = 123000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000.35
	fmt.Println(numReal2)

	// Quando declaramos de forma implicita, o GO se basea na arquitetura do processador, isto é, 32 ou 64bits
	numReal4 := 12345.75
	fmt.Println(numReal4)

	// STRINGS
	var str string = "Text"
	fmt.Println(str)

	str2 := "Text2"
	fmt.Println(str2)

	char := 'B' // aspas simples ele pega o indice da tabela ascii
	fmt.Println(char)

	//VALOR ZERO
	var texto int
	fmt.Println(texto)

	//Boolean
	var boolean1 bool = true // false or empty
	fmt.Println(boolean1)

	//Error
	var erro error = errors.New("Erro interno")
	fmt.Println(erro) //nil -> null or zero

}
