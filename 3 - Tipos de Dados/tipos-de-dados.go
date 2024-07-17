package main

import(
	"errors"
	"fmt"
)

func main()  {
	//int8, int16, int32, int64
	var numero int64 = -1000000000000
	fmt.Println(numero)

	var numero2 uint32 = 100000
	fmt.Println(numero2)

	// alias
	// int32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	//byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12300000.45
	fmt.Println(numeroReal2)
	
	numeroReal3 := 12345.45
	fmt.Println(numeroReal3)

	var str string = "texto"
	fmt.Println(str)

	str2 := "texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	var texto_valor_zero string
	fmt.Println(texto_valor_zero)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)
}
