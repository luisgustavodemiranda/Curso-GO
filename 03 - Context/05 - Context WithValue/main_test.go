package main

import (
	"context"
	"testing"
)

func TestID(t *testing.T) {
	raiz := context.Background()
	if _, ok := idDaRequisicao(raiz); ok {
		t.Fatal("raiz possui ID")
	}
	ctx := comID(raiz, "req-123")
	filho, cancelar := context.WithCancel(ctx)
	defer cancelar()
	if id, ok := idDaRequisicao(filho); !ok || id != "req-123" {
		t.Fatalf("heranca: %q %v", id, ok)
	}
	if _, ok := idDaRequisicao(raiz); ok {
		t.Fatal("contexto pai foi alterado")
	}
	outro := comID(ctx, "req-456")
	if id, _ := idDaRequisicao(outro); id != "req-456" {
		t.Fatal("novo valor ausente")
	}
	if id, _ := idDaRequisicao(ctx); id != "req-123" {
		t.Fatal("valor original alterado")
	}
}
