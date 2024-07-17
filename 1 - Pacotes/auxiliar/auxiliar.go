package auxiliar

import "fmt"

// Escrever registra uma mensagem na tela
func Escrever(){
	fmt.Println("Escrevendo do pacote auxiliar")
	// aqui estou chamando uma função dentro do mesmo pacote dessa onde esse arquivo esta localizado
	escrever2()
}