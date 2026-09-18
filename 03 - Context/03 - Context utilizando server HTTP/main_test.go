package main

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	for _, caso := range []struct {
		nome, url string
		limite    time.Duration
		status    int
		texto     string
	}{
		{"sucesso", "/?rapido=1", time.Second, 200, "Consulta concluida"},
		{"prazo", "/", -time.Second, 504, "Prazo do servidor esgotado"},
		{"rota", "/ausente", time.Second, 404, "404"},
	} {
		t.Run(caso.nome, func(t *testing.T) {
			resp := httptest.NewRecorder()
			handler(caso.limite, time.Hour).ServeHTTP(resp, httptest.NewRequest("GET", caso.url, nil))
			if resp.Code != caso.status || !strings.Contains(resp.Body.String(), caso.texto) {
				t.Fatalf("resposta: %d %s", resp.Code, resp.Body.String())
			}
		})
	}
}
func TestClienteCancelado(t *testing.T) {
	ctx, cancelar := context.WithCancel(context.Background())
	cancelar()
	req := httptest.NewRequest("GET", "/", nil).WithContext(ctx)
	resp := httptest.NewRecorder()
	handler(time.Hour, time.Hour).ServeHTTP(resp, req)
	if resp.Body.Len() != 0 {
		t.Fatal("escreveu resposta para cliente cancelado")
	}
}
