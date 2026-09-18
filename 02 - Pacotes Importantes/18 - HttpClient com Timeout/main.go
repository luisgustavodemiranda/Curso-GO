package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"time"
)

func consultar(cliente *http.Client, url string) error {
	resposta, err := cliente.Get(url)
	if err != nil {
		return err
	}
	defer resposta.Body.Close()
	if resposta.StatusCode != http.StatusOK {
		return fmt.Errorf("status: %s", resposta.Status)
	}
	_, err = io.Copy(io.Discard, resposta.Body)
	return err
}

func main() {
	servidor := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/lento" {
			time.Sleep(300 * time.Millisecond)
		}
		fmt.Fprintln(w, "Resposta local")
	}))
	defer servidor.Close()
	cliente := &http.Client{Timeout: 100 * time.Millisecond}
	if err := consultar(cliente, servidor.URL+"/rapido"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Resposta rapida: OK")
	err := consultar(cliente, servidor.URL+"/lento")
	if erroRede, ok := err.(net.Error); ok && erroRede.Timeout() {
		fmt.Println("Resposta lenta: timeout")
		return
	}
	log.Fatalf("Esperava timeout; recebi: %v", err)
}
