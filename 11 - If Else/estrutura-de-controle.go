package main

import(
	"fmt"
)

func main(){
	numero := 10

	// pode usar parentes no if se tiver uma ordem de precedencia
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("menor ou igual a 15")
	}

	// if init
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("número é maior que zero")
	} 
}