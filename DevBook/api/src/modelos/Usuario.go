package modelos

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"

	"api/src/seguranca"
)

// Usuario representa um usuário da aplicação
type Usuario struct {
	ID           uint64    `json:"id,omitempty"`
	Nome         string    `json:"nome,omitempty"`
	Nick         string    `json:"nick,omitempty"`
	Email        string    `json:"email,omitempty"`
	Senha        string    `json:"senha,omitempty"`
	CriadoEm     time.Time `json:"criadoEm,omitempty"`
	AtualizadoEm time.Time `json:"atualizadoEm,omitempty"`
}

func (usuario *Usuario) Prepara(etapa string) error {
	if erro := usuario.Validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.Formatar(etapa); erro != nil {
		return erro
	}

	return nil
}

func (u *Usuario) Validar(etapa string) error {
	if u.Nome == "" {
		return errors.New("nome é um campo obrigatório")
	}

	if u.Email == "" {
		return errors.New("email é um campo obrigatório")
	}

	if etapa == "cadastro" && u.Senha == "" {
		return errors.New("senha é um campo obrigatório")
	}

	if u.Nick == "" {
		return errors.New("nick é um campo obrigatório")
	}

	if error := checkmail.ValidateFormat(u.Email); error != nil {
		return errors.New("email inválido")
	}
	return nil
}

func (u *Usuario) Formatar(etapa string) error {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(u.Senha)
		if erro != nil {
			return erro
		}

		u.Senha = string(senhaComHash)
	}
	return nil
}
