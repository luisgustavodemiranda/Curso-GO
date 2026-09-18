package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

func main() {
	servidor := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		corpo, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Erro ao ler corpo", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Metodo: %s\nX-Curso: %s\nContent-Type: %s\nCorpo: %s\n",
			r.Method, r.Header.Get("X-Curso"), r.Header.Get("Content-Type"), corpo)
	}))
	defer servidor.Close()
	requisicao, err := http.NewRequest(http.MethodPut, servidor.URL+"/pessoas/1", strings.NewReader(`{"nome":"Maria"}`))
	if err != nil {
		log.Fatal(err)
	}
	requisicao.Header.Set("Content-Type", "application/json")
	requisicao.Header.Set("X-Curso", "Go")
	cliente := &http.Client{Timeout: 3 * time.Second}
	resposta, err := cliente.Do(requisicao)
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
	fmt.Print(string(corpo))
}
