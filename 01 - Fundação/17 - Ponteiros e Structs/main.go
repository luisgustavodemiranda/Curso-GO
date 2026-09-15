package main

import "fmt"

type Pessoa struct {
	Nome string
}

func atualizarNome(pessoa *Pessoa) {
	pessoa.Nome = "Maria"
}

func main() {
	pessoa := Pessoa{Nome: "Joao"}
	atualizarNome(&pessoa)
	fmt.Println(pessoa.Nome)
}
