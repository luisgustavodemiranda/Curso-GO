package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBuscarCEP(t *testing.T) {
	casos := []struct {
		nome, cep, corpo string
		status           int
		erro             error
		falha            bool
	}{
		{"sucesso", "01001000", `{"cep":"01001-000","localidade":"Sao Paulo"}`, 200, nil, false},
		{"formato", "abc", "", 200, errCEPInvalido, true},
		{"inexistente", "99999999", `{"erro":true}`, 200, errCEPInexistente, true},
		{"status", "01001000", "indisponivel", 503, nil, true},
		{"json", "01001000", "{", 200, nil, true},
	}
	for _, caso := range casos {
		t.Run(caso.nome, func(t *testing.T) {
			chamadas := 0
			servidor := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				chamadas++
				w.WriteHeader(caso.status)
				fmt.Fprint(w, caso.corpo)
			}))
			defer servidor.Close()
			endereco, err := buscarCEP(servidor.Client(), servidor.URL, caso.cep)
			if (err != nil) != caso.falha {
				t.Fatalf("erro inesperado: %v", err)
			}
			if caso.erro != nil && !errors.Is(err, caso.erro) {
				t.Fatalf("esperava %v, recebi %v", caso.erro, err)
			}
			if !caso.falha && endereco.CEP != "01001-000" {
				t.Fatalf("endereco: %+v", endereco)
			}
			if caso.nome == "formato" && chamadas != 0 {
				t.Fatal("CEP invalido consultou a API")
			}
		})
	}
	servidor := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	servidor.Close()
	if _, err := buscarCEP(servidor.Client(), servidor.URL, "01001000"); err == nil {
		t.Fatal("falha de transporte ignorada")
	}
}
