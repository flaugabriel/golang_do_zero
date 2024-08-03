package banco

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5" // drive de conexão com o Postgres
)

const PostgresDriver = "postgres"
const User = "root"
const Host = "localhost"
const Port = "5432"
const Password = "root"
const DbName = "devbook"

// conectar abrea conexão com o banco de dados
func Conectar() (*pgx.Conn, error) {
	//var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
	//	"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", "root", "root", "localhost", "5432", "devbook")
	conn, err := pgx.Connect(context.Background(), connStr)

	if err != nil {
		return nil, err
	}
	return conn, nil
}
