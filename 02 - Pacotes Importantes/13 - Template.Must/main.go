package main

import (
	"log"
	"os"
	"text/template"
)

// Must provoca panic se Parse falhar. Aqui o template é conhecido na inicialização.
var modelo = template.Must(template.New("saudacao").Parse("Ola, {{.}}!\n"))

func main() {
	if err := modelo.Execute(os.Stdout, "Ana"); err != nil {
		log.Fatal(err)
	}
}
