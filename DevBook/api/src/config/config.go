package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	StringConexao = ""
	Porta         = "5432"
	WebPort       = "5000"
	SecretKey     = []byte{}
)

func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta = os.Getenv("DB_PORT")
	StringConexao = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USUARIO"), os.Getenv("DB_SENHA"), os.Getenv("DB_HOST"), Porta, os.Getenv("DB_NOME"))
	WebPort = os.Getenv("WEB_PORT")

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
