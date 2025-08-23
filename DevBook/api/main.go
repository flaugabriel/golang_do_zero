package main

import (
	"fmt"
	"log"
	"net/http"

	"api/src/router"

)

func main() {
	fmt.Println("Rodando API")
	router := router.GerarRotas()
	log.Fatal(http.ListenAndServe(":5000", router))
}
