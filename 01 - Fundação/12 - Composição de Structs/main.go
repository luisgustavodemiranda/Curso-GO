package main

import "fmt"

type Endereco struct {
	Cidade string
	Estado string
}

type Pessoa struct {
	Nome string
	Endereco
}

func main() {
	pessoa := Pessoa{
		Nome:     "Joao",
		Endereco: Endereco{Cidade: "Recife", Estado: "PE"},
	}

	fmt.Println(pessoa.Nome, pessoa.Cidade, pessoa.Estado)
}
