package controllers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"

)

// Criar um novo usuário
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Prepara(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close(context.Background())

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	id, erro := repositorio.Criar(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	usuario.ID = id

	respostas.JSON(w, http.StatusCreated, usuario)

}

// Buscar todos os usuários
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := banco.Conectar()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close(context.Background())

	repositorios := repositorios.NovoRepositorioDeUsuarios(db)
	usuarios, erro := repositorios.Buscar(nomeOuNick)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

// Atualizar um usuário
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Usuário atualizado com sucesso"))
	// Lógica para atualizar um usuário
}

// Deletar um usuário
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Usuário deletado com sucesso"))
	// Lógica para deletar um usuário
}

// Buscar um usuário
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Detalhes do usuário"))
	// Lógica para buscar um usuário específico
}
