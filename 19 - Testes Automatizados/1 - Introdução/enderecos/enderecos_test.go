package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

// go test -v - Testa o seu código e mostra alguns dados mais verbosos
// go test ./... - Testa todos os arquivos de todos pacotes
// go test --cover - Mostra a porcentagem de coberta de testes dos pacotes
// go test --coverprofile cobertura.txt - Mostra a porcentagem de cobertura do seu teste naquele pacote e cria um arquivo com os dados de cobertura
// go tool cover --func=cobertura.txt - Mostra a porcentagem e a função
// go tool cover --html=covertura.txt - Mostra um html com seu código mostrando as linhas que não tiveram coberta de teste

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

// TESTE DE UNIDADE
func TestTiposDeEndereco(t *testing.T) {
	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia Dos Imigrantes", "Rodovia"},
		{"Praça das Rodas", "Tipo inválido"},
		{"RUA DOS BOBOS", "Rua"},
		{"Estrada qualquer", "Estrada"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Tipo inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TiposDeEndereco(cenario.enderecoInserido)

		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf(
				"O tipo de endereço recebido é diferente do esperado! Esperava %s e recebeu %s",
				retornoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Error("Teste quebrou!")
	}
}
