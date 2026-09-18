package saudacao

import "strings"

// Mensagem normaliza o nome e cria uma saudação.
func Mensagem(nome string) string {
	return "Ola, " + normalizar(nome) + "!"
}

func normalizar(nome string) string {
	return strings.ToUpper(strings.TrimSpace(nome))
}
