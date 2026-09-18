package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

func buscar(ctx context.Context, cliente *http.Client, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	resp, err := cliente.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status inesperado: %s", resp.Status)
	}
	corpo, err := io.ReadAll(resp.Body)
	return string(corpo), err
}
func executar() error {
	servidor := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/rapido" {
			fmt.Fprint(w, "Resposta recebida")
			return
		}
		timer := time.NewTimer(time.Second)
		defer timer.Stop()
		select {
		case <-r.Context().Done():
			return
		case <-timer.C:
			fmt.Fprint(w, "Resposta lenta")
		}
	}))
	defer servidor.Close()
	cliente := &http.Client{Timeout: 3 * time.Second}
	// O cliente HTTP limita a chamada; cada contexto pode impor um prazo menor.
	rapido, liberarRapido := context.WithTimeout(context.Background(), 2*time.Second)
	defer liberarRapido()
	texto, err := buscar(rapido, cliente, servidor.URL+"/rapido")
	if err != nil {
		return err
	}
	fmt.Println(texto)
	lento, liberarLento := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer liberarLento()
	_, err = buscar(lento, cliente, servidor.URL+"/lento")
	if !errors.Is(err, context.DeadlineExceeded) {
		return fmt.Errorf("esperava prazo esgotado, recebi: %v", err)
	}
	fmt.Println("Requisicao cancelada: prazo do contexto esgotado")
	return nil
}
func main() {
	if err := executar(); err != nil {
		log.Fatal(err)
	}
}
