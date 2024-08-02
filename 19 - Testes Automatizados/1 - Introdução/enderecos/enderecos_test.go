package enderecos_test

import (
	. "introducao-atestes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()

	cenariosDeteste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia ABC", "Rodovia"},
		{"Praça das ABC", "Tipo Inválido"},
		{"Estrada ABC", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA BOBOS", "Avenida"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeteste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s e diferente do esperado %s", retornoRecebido, cenario.retornoEsperado)
		}
	}
}

func TestQualqer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Error("Teste quebrou!")
	}
}
