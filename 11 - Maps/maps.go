package main

import (
	"fmt"
)

func main(){
	usuario := map[string]string {
		"nome" : "Pedro",
		"sobrenome": "silva",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string {
		"nome": {
			"primeiro": "joão",
			"segundo": "pedro",
		},
		"curso": {
			"nome": "Sistemas",
			"campus": "campus 1",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Aquário",
	}

	fmt.Println(usuario2)
}