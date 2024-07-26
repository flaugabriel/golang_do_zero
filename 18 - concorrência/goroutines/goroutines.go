package main

import (
	"fmt"
	"time"
)

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {
	// concorrencia é diferente de paralelismo
	go escrever("olá mundo") // go routine
	escrever("programando em go")
}
