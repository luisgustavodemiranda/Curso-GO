package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestBuscar(t *testing.T) {
	servidor := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/erro" {
			http.Error(w, "indisponivel", 503)
			return
		}
		fmt.Fprint(w, "OK")
	}))
	defer servidor.Close()
	cliente := servidor.Client()
	cliente.Timeout = time.Second
	texto, err := buscar(context.Background(), cliente, servidor.URL)
	if err != nil || texto != "OK" {
		t.Fatalf("sucesso: %q %v", texto, err)
	}
	if _, err := buscar(context.Background(), cliente, servidor.URL+"/erro"); err == nil {
		t.Fatal("status ignorado")
	}
	ctx, cancelar := context.WithCancel(context.Background())
	cancelar()
	if _, err := buscar(ctx, cliente, servidor.URL); !errors.Is(err, context.Canceled) {
		t.Fatalf("cancelamento: %v", err)
	}
}
func TestCancelamentoPropagaAoServidor(t *testing.T) {
	iniciou := make(chan struct{})
	terminou := make(chan struct{})
	servidor := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		close(iniciou)
		<-r.Context().Done()
		close(terminou)
	}))
	defer servidor.Close()
	ctx, cancelar := context.WithCancel(context.Background())
	defer cancelar()
	cliente := servidor.Client()
	cliente.Timeout = 2 * time.Second
	resultado := make(chan error, 1)
	go func() { _, err := buscar(ctx, cliente, servidor.URL); resultado <- err }()
	select {
	case <-iniciou:
	case <-time.After(3 * time.Second):
		t.Fatal("servidor nao recebeu chamada")
	}
	cancelar()
	select {
	case err := <-resultado:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("erro: %v", err)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("cliente nao cancelou")
	}
	select {
	case <-terminou:
	case <-time.After(3 * time.Second):
		t.Fatal("servidor nao observou cancelamento")
	}
}
