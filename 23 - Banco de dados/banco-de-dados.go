package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const PostgresDriver = "postgres"
const User = "root"
const Host = "localhost"
const Port = "5432"
const Password = "root"
const DbName = "devbook"

func main() {
	var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
	db, erro := sql.Open("postgres", DataSourceName)

	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta!")

	linhas, erro := db.Query("select * from usuarios")

	if erro != nil {
		log.Fatal(erro)
	}

	defer linhas.Close()

	fmt.Println(linhas)
}
