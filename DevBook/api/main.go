package main

import (
	"fmt"
	"log"
	"net/http"

	"api/src/config"
	"api/src/router"
)

// func init() {
// 	chave := make([]byte, 64)
// 	if _, erro := rand.Read(chave); erro != nil {
// 		log.Fatal(erro)
// 	}

// 	stringBase64 := base64.StdEncoding.EncodeToString(chave)
// 	fmt.Println("Chave de autenticação gerada:", stringBase64)
// }

func main() {
	config.Carregar()
	fmt.Println("localhost:" + config.WebPort)
	fmt.Println("Rodando API")
	r := router.GerarRotas()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.WebPort), r))
}
