package main

import(
	"fmt"
)

func main(){
	totalDaSoma := soma(1,2,23,44,35,6)
	fmt.Println(totalDaSoma)
	escrever("ola mundo", 12,51,15,44)
}

// aqui ela vai receber de 1 a N int's
func soma(numeros ...int) int{
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int){
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}