package modelos

import (
	"errors"
	"strings"
	"time"
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

func (usuario *Usuario) Prepara() error {
	if erro := usuario.Validar(); erro != nil {
		return erro
	}

	usuario.Formatar()
	return nil
}

func (u *Usuario) Validar() error {
	if u.Nome == "" {
		return errors.New("nome é um campo obrigatório")
	}

	if u.Email == "" {
		return errors.New("email é um campo obrigatório")
	}

	if u.Senha == "" {
		return errors.New("senha é um campo obrigatório")
	}

	if u.Nick == "" {
		return errors.New("nick é um campo obrigatório")
	}
	return nil
}

func (u *Usuario) Formatar() error {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
	return nil
}
