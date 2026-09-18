package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

func main() {
	servidor := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-time.After(time.Second):
			fmt.Fprintln(w, "Resposta demorada")
		case <-r.Context().Done():
			// O cliente cancelou; não há mais motivo para continuar o trabalho.
			return
		}
	}))
	defer servidor.Close()
	ctx, cancelar := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancelar()
	requisicao, err := http.NewRequestWithContext(ctx, http.MethodGet, servidor.URL, nil)
	if err != nil {
		log.Fatal(err)
	}
	cliente := &http.Client{Timeout: 3 * time.Second}
	resposta, err := cliente.Do(requisicao)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("Requisicao cancelada: prazo do contexto esgotado")
			return
		}
		log.Fatal(err)
	}
	defer resposta.Body.Close()
	log.Fatal("Esperava cancelamento, mas a requisicao terminou")
}
