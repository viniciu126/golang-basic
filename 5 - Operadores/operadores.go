package main

import "fmt"

func main() {
	// ARITMETICOS
	soma := 1 + 2
	subtracao := 3 - 2
	divisao := 2 / 2
	multiplicacao := 1 * 2
	restoDaDivisao := 10 % 2

	fmt.Println(soma)
	fmt.Println(subtracao)
	fmt.Println(divisao)
	fmt.Println(multiplicacao)
	fmt.Println(restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int32 = 25

	soma2 := numero1 + int16(numero2)
	fmt.Println(soma2)
	// FIM ARITMETICOS

	// ATRIBUIÇÃO
	var variavel string = "String"
	variavel2 := "String2"
	fmt.Println(variavel, variavel2)
	// FIM ATRIBUIÇÃO

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	// FIM DOS RELACIONAIS

	// OPERADORES LÓGICOS
	fmt.Println("----------------")
	fmt.Println(1 < 2 && true)
	fmt.Println(1 < 3 || true)
	fmt.Println(!true)
	fmt.Println(!false)
	// FIM DOS OPERADORES LÓGICAS

	// OPERADORES UNÁRIOS
	numero := 10
	numero++
	numero += 15
	numero--
	numero -= 15
	numero *= 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)
	// FIM OPERADORES UNÁRIOS

	// NÃO TEM OPERADOR TERNÁRIO, USE IF E ELSE MESMO
}
