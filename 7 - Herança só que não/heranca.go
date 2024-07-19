package main

import(
	"fmt"
)

func main(){
	p1 := pessoa{"joão", "Pedro", 20, 178}
	fmt.Println(p1)
	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1.nome)
}

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}