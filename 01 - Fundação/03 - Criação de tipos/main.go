package main

import "fmt"

type Usuario struct {
	Nome  string
	Ativo bool
}

type ID int

func main() {
	usuario := Usuario{Nome: "Ana", Ativo: true}
	var codigo ID = 42

	fmt.Println(usuario.Nome, usuario.Ativo)
	fmt.Println(codigo)
}
