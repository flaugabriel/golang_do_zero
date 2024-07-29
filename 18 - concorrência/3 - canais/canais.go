package main

import (
	"fmt"
	"time"
)

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}

func main() {
	canal := make(chan string)

	go escrever("olá mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	//for {
	//	mensagem, aberto := <-canal //passando o segundo atributo "aberto" e possivel verifica se o canal esta aberto ou fechado = true/false
	//	if !aberto {
	//		break // maneira de sair de um loop infinito
	//	}
	//	fmt.Println(mensagem)
	//}

	fmt.Println("Fim do programa!")
}
