package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

func main() {
	// Este servidor local permite estudar sem depender da internet.
	servidor := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ola do servidor HTTP!")
	}))
	defer servidor.Close()

	cliente := &http.Client{Timeout: 3 * time.Second}
	resposta, err := cliente.Get(servidor.URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resposta.Body.Close()
	if resposta.StatusCode != http.StatusOK {
		log.Fatalf("Status inesperado: %s", resposta.Status)
	}
	corpo, err := io.ReadAll(resposta.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Status:", resposta.Status)
	fmt.Print(string(corpo))
}
