package enderecos

import "strings"

// TipoEndereco verifica se o endereço é válido e o retorna
func TipoEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoMinusculo := strings.ToLower(endereco)
	primeiraPalavra := strings.Split(enderecoMinusculo, " ")[0]

	enderecoTemTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			enderecoTemTipoValido = true
		}
	}

	if enderecoTemTipoValido {
		return strings.Title(primeiraPalavra)
	}

	return "Tipo Inválido"

}
