package main

import "fmt"

func main() {
	defer recuperaExecucao()
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("pós execução")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é exatamente 6!")
}

func recuperaExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Tentando recupera a função")
	}
}
