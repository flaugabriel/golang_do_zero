package seguranca

import (
	"golang.org/x/crypto/bcrypt"

)

// Hash gera um hash para a senha recebida
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}


func VerificarSenha(senhaHash, senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senha))
}

