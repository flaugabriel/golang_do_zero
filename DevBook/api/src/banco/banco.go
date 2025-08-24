package banco

import (
	"context"

	"github.com/jackc/pgx/v5" // drive de conexão com o Postgres

	"api/src/config"
)

// conectar abrea conexão com o banco de dados
func Conectar() (*pgx.Conn, error) {
	config.Carregar()
	conn, err := pgx.Connect(context.Background(), config.StringConexao)

	if err != nil {
		return nil, err
	}

	if err = conn.Ping(context.Background()); err != nil {
		conn.Close(context.Background())
		return nil, err

	}

	return conn, nil
}
