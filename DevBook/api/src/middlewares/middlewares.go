package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"api/src/autenticacao"
)

// Logger é um middleware que loga as requisições recebidas
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Requisição recebida: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

// Autenticar é um middleware que autentica as requisições
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			http.Error(w, fmt.Sprintf("Não autorizado: %v", erro), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
