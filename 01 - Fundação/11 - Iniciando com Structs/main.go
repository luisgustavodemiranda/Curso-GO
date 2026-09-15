package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	pessoa := Pessoa{Nome: "Joao", Idade: 30}
	fmt.Printf("%s tem %d anos.\n", pessoa.Nome, pessoa.Idade)
}
