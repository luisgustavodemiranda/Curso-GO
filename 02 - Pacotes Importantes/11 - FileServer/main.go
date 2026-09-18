package main

import (
	"log"
	"net/http"
	"time"
)

func rotas() http.Handler {
	// Execute a partir da pasta desta aula: somente public será exposta.
	return http.FileServer(http.Dir("public"))
}
func main() {
	servidor := &http.Server{
		Addr:              "localhost:8080",
		Handler:           rotas(),
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Println("Acesse http://localhost:8080 — encerre com Ctrl+C")
	log.Fatal(servidor.ListenAndServe())
}
