package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

func main() {
	original := Pessoa{Nome: "Ana", Idade: 25}
	dados, err := json.Marshal(original)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("JSON:", string(dados))

	var pessoa Pessoa
	if err := json.Unmarshal(dados, &pessoa); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Pessoa: %s, %d anos\n", pessoa.Nome, pessoa.Idade)
}
