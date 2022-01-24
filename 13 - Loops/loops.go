package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		time.Sleep(time.Second)
		fmt.Println("Incrementando i")
		i++
	}

	fmt.Println(i)

	for j := 0; j < 10; j += 2 {
		time.Sleep(time.Second)
		fmt.Println("Incrementando j", j)
	}

	nomes := []string{"João", "Davi", "Vini"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}
}
