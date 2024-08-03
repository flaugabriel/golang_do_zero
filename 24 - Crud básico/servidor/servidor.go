package servidor

import (
	"context"
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type usuario struct {
	Id    uint   `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario inseri um usuário ao db
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao converter conecta no banco de dados!"))
		return
	}
	defer db.Close(context.Background())

	// prepare statement não sei se utiliza no pg usando go mais adaptei para pg?
	query := `INSERT INTO public.usuarios (nome, email) VALUES($1, $2) RETURNING  id`
	data, erro := db.Exec(context.Background(), query, usuario.Nome, usuario.Email)

	if erro != nil {
		fmt.Println(erro)
		log.Println("erro ao inserir usuário")
		return
	}
	fmt.Println(data.Insert())
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(("Usuário inserido com sucesso!")))
}
