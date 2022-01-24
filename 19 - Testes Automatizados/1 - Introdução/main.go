package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TiposDeEndereco("Rodovia dos imigrantes")
	fmt.Println(tipoEndereco)
}
