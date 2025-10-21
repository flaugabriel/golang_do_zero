package modelos

import (
	"errors"
	"strings"
	"time"

)

type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autorId,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadaEm  time.Time `json:"criadaEm,omitempty"`
}

func (publicacao *Publicacao) Preparar() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}

	publicacao.formatar()
	return nil
}

func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("o titulo e obrigatorio e nao pdoe esta em branco")
	}

	if publicacao.Conteudo == "" {
		return errors.New("o conteudo e obrigatorio e nao pode esta em branco")
	}

	return nil
}

func (Publicacao *Publicacao) formatar() {
	Publicacao.Titulo = strings.TrimSpace(Publicacao.Titulo)
	Publicacao.Conteudo = strings.TrimSpace(Publicacao.Conteudo)
}
