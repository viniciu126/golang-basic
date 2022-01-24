package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá mundo", canal)

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")
}

func escrever(texto string, canal chan string) {
	defer close(canal)

	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
}
