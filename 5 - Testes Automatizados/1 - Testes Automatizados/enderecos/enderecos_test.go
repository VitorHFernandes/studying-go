// TESTE DE UNIDADE
package enderecos

import (
	"testing"
	"time"
)

type cenarioTeste struct {
	enderecoInserido string
	enderecoEsperado string
}

func TestTipoEndereco(t *testing.T) {
	t.Parallel() // testa as funções em paralelo
	// TestXxxxx - nome da função
	cenariosDeTeste := []cenarioTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Brasil", "Avenida"},
		{"Rodovia do Arroz", "Rodovia"},
		{"Estrada Timbé", "Estrada"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"RUA DOS BOBOS", "Rua"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoEnderecoRecebido := TipoEndereco(cenario.enderecoInserido)

		if tipoEnderecoRecebido != cenario.enderecoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperado '%s', e o recebido foi '%s'", tipoEnderecoRecebido, cenario.enderecoEsperado)
		}
	}

}
func TestQualquer(t *testing.T) {
	t.Parallel() // testa as funções em paralelo

	time.Sleep(time.Second * 5)
	if 1 > 2 {
		t.Errorf("Infelizmente, neste universo %d é maior que %d", 1, 2)
	}
}
