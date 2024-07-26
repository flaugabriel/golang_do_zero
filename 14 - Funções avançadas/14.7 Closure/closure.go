package main

import "fmt"

func main() {
	texto := "dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}

func closure() func() {
	texto := "Dentro da função closure"
	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}
