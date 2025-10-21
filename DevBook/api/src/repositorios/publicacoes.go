package repositorios

import (
	"context"

	"github.com/jackc/pgx/v5"

	"api/src/modelos"
)

type Publicacoes struct {
	db *pgx.Conn
}

func NovoRepositorioDePublicacoes(db *pgx.Conn) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorios Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	var id uint64
	err := repositorios.db.QueryRow(
		context.Background(),
		"INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES ($1, $2, $3) RETURNING id",
		publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(
		context.Background(),
		`SELECT p.*, u.nick FROM publicacoes p INNER JOIN usuarios u ON u.id = p.autor_id WHERE p.id = $1`,
		publicacaoID,
	)

	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linhas.Close()

	var publicacao modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao
		if erro := linhas.Scan(&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo, &publicacao.AutorID, &publicacao.Curtidas, &publicacao.CriadaEm, &publicacao.AutorNick); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil
}
