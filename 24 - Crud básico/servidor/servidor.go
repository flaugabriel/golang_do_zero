package servidor

import (
	"context"
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	Id    string `json:"id"`
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

	fmt.Print(erro)
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

// BuscarUsuarios busca todos os usuários no db
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close(context.Background())

	linhas, erro := db.Query(context.Background(), "SELECT * FROM public.usuarios")

	if erro != nil {
		w.Write([]byte("Erro ao buscar usuários!"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario

	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.Id, &usuario.Nome, &usuario.Email); erro != nil {
			fmt.Print(erro)
			w.Write([]byte("Erro ao escanear usuário!"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter usuários para JSON!"))
		return
	}
}

// BuscarUsuario busca um usuário no db
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	id, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o ID do usuário!"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close(context.Background())

	var usuario usuario
	linha := db.QueryRow(context.Background(), "SELECT * FROM public.usuarios WHERE id = $1", id)
	if erro := linha.Scan(&usuario.Id, &usuario.Nome, &usuario.Email); erro != nil {
		w.Write([]byte("Erro ao buscar usuário!"))
		return
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuário para JSON!"))
		return
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	id, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o ID do usuário!"))
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close(context.Background())

	_, erro = db.Exec(context.Background(), `UPDATE public.usuarios SET nome = $1, email = $2 WHERE id = $3`, usuario.Nome, usuario.Email, id)
	if erro != nil {
		w.Write([]byte("Erro ao atualizar usuário!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	id, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o ID do usuário!"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close(context.Background())

	_, erro = db.Exec(context.Background(), `DELETE FROM public.usuarios WHERE id = $1`, id)
	if erro != nil {
		w.Write([]byte("Erro ao deletar usuário!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
