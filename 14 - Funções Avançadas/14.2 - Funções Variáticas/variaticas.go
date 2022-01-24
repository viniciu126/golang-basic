package main

import "fmt"

func soma(numeros ...int) (total int) {
	fmt.Println(numeros)
	for _, numero := range numeros {
		total += numero
	}

	return
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(10, 20, 30)
	fmt.Println(totalDaSoma)

	escrever("Olá Mundo", 10, 20, 1)
}
