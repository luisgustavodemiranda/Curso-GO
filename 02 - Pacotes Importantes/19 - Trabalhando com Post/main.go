package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

type Pessoa struct {
	Nome string `json:"nome"`
}

func receberPessoa(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "use POST", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	var pessoa Pessoa
	if err := json.NewDecoder(r.Body).Decode(&pessoa); err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(pessoa); err != nil {
		log.Println(err)
	}
}

func main() {
	servidor := httptest.NewServer(http.HandlerFunc(receberPessoa))
	defer servidor.Close()
	dados, err := json.Marshal(Pessoa{Nome: "Ana"})
	if err != nil {
		log.Fatal(err)
	}
	cliente := &http.Client{Timeout: 3 * time.Second}
	resposta, err := cliente.Post(servidor.URL, "application/json", bytes.NewReader(dados))
	if err != nil {
		log.Fatal(err)
	}
	defer resposta.Body.Close()
	if resposta.StatusCode != http.StatusCreated {
		log.Fatalf("Status inesperado: %s", resposta.Status)
	}
	corpo, err := io.ReadAll(resposta.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Status:", resposta.Status)
	fmt.Print(string(corpo))
}
