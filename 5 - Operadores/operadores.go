package main

import(
	"fmt"
)

func main()  {
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 2
	multiplicacao := 10 * 25
	restodadivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restodadivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1  + numero2
	fmt.Println(soma2)

	// atribuiçao
	var variavel1 string = "String texto"
	variavel2 := "string 2"
	fmt.Println(variavel1, variavel2)

	// operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	
	//operadores logicos
	fmt.Println("-----------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// operadores unários
	numero := 10
	numero = numero + 1
	numero++
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--
	numero-=20
	fmt.Println(numero)
	numero *= 3 // numero = numero * 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)

	// operador ternario
	// no go so pode existr uma maneira de fazer as coisas
	// Errado: texto := numero > 5 ? "Maior que 5" : "Menor que 5"
	var texto string
	
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
}