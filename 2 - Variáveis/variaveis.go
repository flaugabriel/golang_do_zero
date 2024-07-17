package main

import "fmt"

func main()  {
	var variavel1 string  = "Variável 1"
	variavel2 := "Variável 2" // Declaração implicita quando a variavel e declarada mais ainda não sabe qual é o tipo dela

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "lalala"
		variavel4 string = "tralalal"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável 5", "variavel 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	//trocando os valores
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}