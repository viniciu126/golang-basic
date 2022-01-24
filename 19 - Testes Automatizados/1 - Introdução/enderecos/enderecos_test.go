package enderecos

import "testing"

// TESTE DE UNIDADE
func TestTiposDeEndereco(t *testing.T) {
	enderecoParaTeste := "Avenida paulista"
	tipoDeEnderecoEsperado := "Avenida"
	tipoDeEnderecpRecebido := TiposDeEndereco(enderecoParaTeste)

	if tipoDeEnderecpRecebido != tipoDeEnderecoEsperado {
		t.Error("O tipo de endereço recebido é diferente do esperado!")
	}
}
