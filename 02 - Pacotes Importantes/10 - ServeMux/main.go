package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func rotas() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// O padrão / também captura caminhos sem rota mais específica.
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintln(w, "Pagina inicial")
	})
	mux.HandleFunc("/sobre", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Curso de Go — Pacotes Importantes")
	})
	return mux
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
