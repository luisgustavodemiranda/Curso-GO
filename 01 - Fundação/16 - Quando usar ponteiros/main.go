package main

import "fmt"

type Configuracao struct {
	Nome string
}

func atualizar(config *Configuracao) {
	config.Nome = "producao"
}

func main() {
	config := Configuracao{Nome: "teste"}
	atualizar(&config)
	fmt.Println(config.Nome)
}
