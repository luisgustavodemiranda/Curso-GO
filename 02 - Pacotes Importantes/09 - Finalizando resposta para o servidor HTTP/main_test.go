package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestRespostaHTTP(t *testing.T) {
	servidor := servidorLocal()
	defer servidor.Close()
	handler := rotas(servidor.Client(), servidor.URL)
	casos := []struct {
		metodo, caminho string
		status          int
	}{
		{"GET", "/cep?cep=01001000", 200},
		{"GET", "/cep", 400},
		{"GET", "/cep?cep=123", 400},
		{"GET", "/cep?cep=99999999", 404},
		{"POST", "/cep?cep=01001000", 405},
	}
	for _, caso := range casos {
		req := httptest.NewRequest(caso.metodo, caso.caminho, nil)
		resp := httptest.NewRecorder()
		handler.ServeHTTP(resp, req)
		if resp.Code != caso.status || !json.Valid(resp.Body.Bytes()) {
			t.Fatalf("%s: resposta inesperada: %d %s", caso.caminho, resp.Code, resp.Body.String())
		}
		if !strings.HasPrefix(resp.Header().Get("Content-Type"), "application/json") {
			t.Fatal("Content-Type incorreto")
		}
	}
	servidor.Close()
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, httptest.NewRequest("GET", "/cep?cep=01001000", nil))
	if resp.Code != 502 {
		t.Fatalf("API indisponivel: status %d", resp.Code)
	}
}
