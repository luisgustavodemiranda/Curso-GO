package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	funcoes := template.FuncMap{"maiusculas": strings.ToUpper}
	// Registre as funções antes de analisar o template.
	modelo, err := template.New("saudacao").Funcs(funcoes).Parse("Ola, {{maiusculas .}}!\n")
	if err != nil {
		log.Fatal(err)
	}
	if err := modelo.Execute(os.Stdout, "Ana"); err != nil {
		log.Fatal(err)
	}
}
