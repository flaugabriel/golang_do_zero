package main

import (
	"fmt"
)

func main(){


	func (){
		fmt.Println("Olá mundo")
	}()

	func (parametro string){
		fmt.Println(parametro)
	}("Passando Parámetro string")

	retorno := func (parametro string) string{
		return fmt.Sprintf("Recebido -> %s", parametro)
	}("Passando Parámetro string")

	fmt.Println(retorno)
}