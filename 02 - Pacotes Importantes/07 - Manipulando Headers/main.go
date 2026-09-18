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
		// Headers precisam ser definidos antes de WriteHeader ou Write.
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Curso", "Go")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "User-Agent:", r.Header.Get("User-Agent"))
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
