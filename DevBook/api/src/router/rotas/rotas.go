package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa uma rota da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configura(r *mux.Router) *mux.Router {
	rotas := RotasUsuarios

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}
