package main

import (
	"fmt"
	"net/http"
)

func main() {
	requisicao, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println(requisicao.Method, requisicao.URL.Host)
}
