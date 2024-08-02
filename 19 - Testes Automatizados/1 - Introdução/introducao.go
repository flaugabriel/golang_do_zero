package main

import (
	"fmt"
	"introducao-atestes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Testes altomatizados")
	fmt.Println(tipoEndereco)
}
