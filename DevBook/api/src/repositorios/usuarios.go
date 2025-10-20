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

func (repositorio Usuarios) Atualizar(id uint64, usuario modelos.Usuario) error {
	_, err := repositorio.db.Exec(
		context.Background(),
		"UPDATE usuarios SET nome = $1, nick = $2, email = $3 WHERE id = $4",
		usuario.Nome, usuario.Nick, usuario.Email, id,
	)

	if err != nil {
		return err
	}
	defer repositorio.db.Close(context.Background())

	return nil
}

func (repositorio Usuarios) Deletar(id uint64) error {
	_, err := repositorio.db.Exec(
		context.Background(),
		"DELETE FROM usuarios WHERE id = $1",
		id,
	)

	if err != nil {
		return err
	}
	defer repositorio.db.Close(context.Background())

	return nil
}

func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linhas, err := repositorio.db.Query(
		context.Background(),
		"SELECT id, senha FROM usuarios WHERE email = $1",
		email,
	)

	if err != nil {
		return modelos.Usuario{}, err
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro := linhas.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repositorio Usuarios) Seguir(usuarioID, seguidorID uint64) error {
	_, erro := repositorio.db.Exec(
		context.Background(),
		"INSERT INTO seguidores (usuario_id, seguidor_id) VALUES ($1, $2) ON CONFLICT DO NOTHING",
		usuarioID, seguidorID,
	)

	if erro != nil {
		return erro
	}

	defer repositorio.db.Close(context.Background())

	return nil
}

// PararDeSeguir permite que um usuário pare de seguir outro usuário
func (repositorio Usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {
	_, erro := repositorio.db.Exec(
		context.Background(),
		"DELETE FROM seguidores WHERE usuario_id = $1 AND seguidor_id = $2",
		usuarioID, seguidorID,
	)
	if erro != nil {
		return erro
	}

	defer repositorio.db.Close(context.Background())

	return nil
}

func (repositorio Usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		context.Background(),
		`SELECT u.id, u.nome, u.nick, u.email, u.criado_em
		FROM usuarios u
		INNER JOIN seguidores s ON u.id = s.seguidor_id
		WHERE s.usuario_id = $1`,
		usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var seguidores []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}
		seguidores = append(seguidores, usuario)
	}

	return seguidores, nil
}

func (repositorio Usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		context.Background(),
		`SELECT u.id, u.nome, u.nick, u.email, u.criado_em
		FROM usuarios u
		INNER JOIN seguidores s ON u.id = s.usuario_id	
		WHERE s.seguidor_id = $1`,
		usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var seguindo []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}
		seguindo = append(seguindo, usuario)
	}
	return seguindo, nil
}

func (repositorio Usuarios) AtualizarSenha(usuarioID uint64, novaSenha string) error {
	_, erro := repositorio.db.Exec(
		context.Background(),
		"UPDATE usuarios SET senha = $1 WHERE id = $2",
		novaSenha, usuarioID,
	)

	if erro != nil {
		return erro
	}

	defer repositorio.db.Close(context.Background())

	return nil
}

func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, erro := repositorio.db.Query(
		context.Background(),
		"SELECT senha FROM usuarios WHERE id = $1",
		usuarioID,
	)

	if erro != nil {
		return "", erro
	}
	defer linha.Close()
	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}
	return usuario.Senha, nil
}
