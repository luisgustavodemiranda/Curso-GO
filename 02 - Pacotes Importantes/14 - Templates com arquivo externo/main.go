package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	modelo, err := template.ParseFiles("templates/mensagem.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	pessoa := struct{ Nome string }{Nome: "Ana"}
	if err := modelo.Execute(os.Stdout, pessoa); err != nil {
		log.Fatal(err)
	}
}
