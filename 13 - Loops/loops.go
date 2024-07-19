package main

import(
	"fmt"
	//"time"
)

func main() {
	//i := 0

//	for i < 10 {
		//fmt.Println("incrementando i")
	//	time.Sleep(time.Second)
//	}

	//fmt.Println(i)
	//for j := 0; j < 10; j++ {
		//fmt.Println("Incrementando j", j)
		//time.Sleep(time.Second)
//	}

//	fmt.Println("fim variavel dentro do for")

	nomes := [3]string{"joao", "savio", "lucas"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string {
		"nome": "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
	
	// loop em structs não se faz.
	//type usuarioStruct struct {
	//	nome string
	//	sobrenome string
	//}

	//usuario2 := usuarioStruct{"zé", "júnior"}

	//for chave, valor := range usuario2 {
		fmt.Println(chave, valor)
	//}

	// loop infinito
	for {
		time.Sleep(time.Second)
	}
}