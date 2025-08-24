package main

import (
	"fmt"
	"log"
	"net/http"

	"api/src/config"
	"api/src/router"

)

func main() {
	config.Carregar()
	fmt.Println("localhost:" + config.WebPort)
	fmt.Println("Rodando API")
	r := router.GerarRotas()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.WebPort), r))
}
