package main

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestReceberPessoa(t *testing.T) {
	casos := []struct {
		metodo, corpo string
		status        int
	}{
		{"POST", `{"nome":"Ana"}`, 201},
		{"POST", "{", 400},
		{"GET", "", 405},
	}
	for _, caso := range casos {
		req := httptest.NewRequest(caso.metodo, "/", strings.NewReader(caso.corpo))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		receberPessoa(resp, req)
		if resp.Code != caso.status {
			t.Fatalf("status: %d", resp.Code)
		}
		if caso.status == 201 {
			var pessoa Pessoa
			if err := json.Unmarshal(resp.Body.Bytes(), &pessoa); err != nil || pessoa.Nome != "Ana" {
				t.Fatalf("resposta: %s", resp.Body.String())
			}
		}
	}
}
