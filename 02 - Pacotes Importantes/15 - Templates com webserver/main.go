package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"time"
)

func rotas() http.Handler {
	// html/template escapa os dados inseridos no HTML.
	modelo := template.Must(template.ParseFiles("templates/pagina.html"))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		nome := r.URL.Query().Get("nome")
		if nome == "" {
			nome = "Ana"
		}
		var pagina bytes.Buffer
		// Renderiza antes de enviar para poder retornar 500 se houver erro.
		if err := modelo.Execute(&pagina, struct{ Nome string }{nome}); err != nil {
			http.Error(w, "Erro ao renderizar pagina", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if _, err := pagina.WriteTo(w); err != nil {
			log.Println(err)
		}
	})
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
