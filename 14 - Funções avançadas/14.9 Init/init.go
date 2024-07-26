package main

import "fmt"

// declarando uma variavel global
var n int

func main() {
	fmt.Println("função main sendo executada")
	fmt.Println(n)
}

func init() {
	fmt.Println(("Executando a função init"))
	n = 10
}
