package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	modelo, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	dados := struct{ Titulo, Nome string }{"Curso de Go", "Ana"}
	if err := modelo.ExecuteTemplate(os.Stdout, "pagina", dados); err != nil {
		log.Fatal(err)
	}
}
