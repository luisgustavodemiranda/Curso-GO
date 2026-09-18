package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

func consultar(ctx context.Context, atraso time.Duration) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	timer := time.NewTimer(atraso)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
func handler(limite, atraso time.Duration) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		// Derivar de r.Context preserva o cancelamento vindo do cliente.
		ctx, cancelar := context.WithTimeout(r.Context(), limite)
		defer cancelar()
		demora := atraso
		if r.URL.Query().Get("rapido") == "1" {
			demora = 0
		}
		if err := consultar(ctx, demora); err != nil {
			if r.Context().Err() != nil {
				log.Println("Cliente cancelou a requisicao")
				return // O cliente pode já ter desconectado.
			}
			if errors.Is(err, context.DeadlineExceeded) {
				http.Error(w, "Prazo do servidor esgotado", http.StatusGatewayTimeout)
				return
			}
			http.Error(w, "Erro na consulta", http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "Consulta concluida")
	})
}
func main() {
	servidor := &http.Server{Addr: "localhost:8080", Handler: handler(100*time.Millisecond, time.Second), ReadHeaderTimeout: 5 * time.Second}
	log.Println("Acesse http://localhost:8080 — encerre com Ctrl+C")
	log.Fatal(servidor.ListenAndServe())
}
