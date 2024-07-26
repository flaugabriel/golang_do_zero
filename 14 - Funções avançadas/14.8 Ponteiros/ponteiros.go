package main

import "fmt"

func main() {
	// não mecher no valor davariavel numero
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	// meche no valor da variavel
	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}
