package main

import "fmt"

func closure() func() {
	texto := "Dentro da função Closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da função main"

	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
