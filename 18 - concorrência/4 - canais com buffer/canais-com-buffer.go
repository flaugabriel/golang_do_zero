package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "Olá mundo!"
	canal <- "Programando em go!"
	// canal <- "Programando em go!" dead lock por que o limite de buff é 2

	mensagem := <-canal
	mensagem2 := <-canal
	//mensagem3 := <-canal dead lock
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
	//fmt.Println(mensagem3) dead lock
}
