package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRotas(t *testing.T) {
	handler := rotas()
	casos := []struct {
		metodo, caminho string
		status          int
		conteudo        string
	}{
		{"GET", "/", 200, "FileServer em Go"},
		{"GET", "/main.go", 404, "404"},
	}
	for _, caso := range casos {
		t.Run(caso.caminho, func(t *testing.T) {
			req := httptest.NewRequest(caso.metodo, caso.caminho, nil)
			req.Header.Set("User-Agent", "EstudanteGo")
			resp := httptest.NewRecorder()
			handler.ServeHTTP(resp, req)
			if resp.Code != caso.status || !strings.Contains(resp.Body.String(), caso.conteudo) {
				t.Fatalf("resposta inesperada: %d %s", resp.Code, resp.Body.String())
			}

		})
	}
}
