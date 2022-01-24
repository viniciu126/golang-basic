package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		/**
		O Select é parecido com switch, porém utilizado somente para concorrência
		É uma estrutura de controle que faz o tempo de recebimento de uma mensagem
		de um canal não depender do outro.

		Por exemplo: O canal 1 envia mensagens a cada 500 milisegundos e o canal 2
		a cada 2 segundos, quando utilizamos o Select as mensagens serão printadas
		de acordo com o time.Sleep definido a cada publicação, indiferente do canal;
		*/
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}
}
