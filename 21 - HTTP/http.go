package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("olá mundo!"))
}

func main() {
	// HTTP é um protocolo de comunicação é a base da comunicação web
	// cliente (faz a requisição) x servidor (processa requisição e envia a resposta)
	// request - response
	// Rotas
	// URI - identificadaor do recurso
	// método - GET POST DELETE PUT
	http.HandleFunc("/home", home)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
