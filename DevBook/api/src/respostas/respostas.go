package respostas

import (
	"encoding/json"
	"log"
	"net/http"

)

// JSON responde com um JSON e o status HTTP
func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)

	if erro := json.NewEncoder(w).Encode(data); erro != nil {
		log.Fatal(erro)
	}

}

func Erro(w http.ResponseWriter, status int, erro error) {
	JSON(w, status, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
