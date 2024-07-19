package main

import(
	"fmt"
)

type usuario struct {
	nome string
	idade uint8
	endereco endereco

}

type endereco struct {
	logradouro string
	numero uint8
}

func main()  {
	fmt.Println("Arquivo structs")

	var u usuario
	fmt.Println(u)
	u.nome = "Samuel"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos bobos", 0}
	fmt.Println(enderecoExemplo)

	usuario2 := usuario{"Davi", 21, enderecoExemplo}
	fmt.Println(usuario2) 
	
	usuario3 := usuario{nome: "Julía"}
	fmt.Println(usuario3)
}