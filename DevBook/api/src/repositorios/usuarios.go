package repositorios

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5" // driver de conexão com o Postgres

	"api/src/modelos"

)

type Usuarios struct {
	db *pgx.Conn
}

// NovoRepositorioDeUsuarios cria um novo repositório de usuários
func NovoRepositorioDeUsuarios(db *pgx.Conn) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	var id uint64
	err := repositorio.db.QueryRow(
		context.Background(),
		"INSERT INTO usuarios (nome, nick, email, senha) VALUES ($1, $2, $3, $4) RETURNING id",
		usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%
	linhas, erro := repositorio.db.Query(
		context.Background(),
		"select id, nome, nick, email, criado_em from usuarios where nome LIKE $1 or nick LIKE $2",
		nomeOuNick, nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio Usuarios) BuscarPorID(id uint64) (modelos.Usuario, error) {
	linhas, err := repositorio.db.Query(
		context.Background(),
		"SELECT id, nome, nick, email, criado_em FROM usuarios WHERE id = $1",
		id,
	)

	if err != nil {
		return modelos.Usuario{}, err
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}
