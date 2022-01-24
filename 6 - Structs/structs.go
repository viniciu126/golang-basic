package main

import "fmt"

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos bobos", 0}

	usuario2 := usuario{"Davi", 21, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Davi"}
	fmt.Println(usuario3)
}
