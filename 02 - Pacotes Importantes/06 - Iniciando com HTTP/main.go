package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func rotas() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintln(w, "Ola, HTTP!")
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
