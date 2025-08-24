package router

import (
	"github.com/gorilla/mux"

	"api/src/router/rotas"

)

	
func GerarRotas() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configura(r)
}
