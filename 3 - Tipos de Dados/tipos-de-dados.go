package main

import (
	"errors"
	"fmt"
)

func main() {
	// aceita negativo
	var numero int = 10000000
	fmt.Println(numero)

	// uint não aceita negativo
	var numero2 uint64 = 100000
	fmt.Println(numero2)

	// alias
	// int32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	// Byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 1.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1252312313.45
	fmt.Println(numeroReal2)

	// FIM NUMEROS REAIS

	var str string = "TEXTO"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	char := '%'
	fmt.Println(char)

	// FIM STRINGS

	var texto string
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
