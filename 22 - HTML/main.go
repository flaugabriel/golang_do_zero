package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))
	// HTTP é um protocolo de comunicação é a base da comunicação web
	// cliente (faz a requisição) x servidor (processa requisição e envia a resposta)
	// request - response
	// Rotas
	// URI - identificadaor do recurso
	// método - GET POST DELETE PUT
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{
			"joão",
			"joao.pedro@gmail.com",
		}
		templates.ExecuteTemplate(w, "home.html", u)
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
