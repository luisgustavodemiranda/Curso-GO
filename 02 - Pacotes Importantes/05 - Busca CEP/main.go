package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"time"
)

type Endereco struct {
	CEP        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
	Erro       bool   `json:"erro,omitempty"`
}

// Simula apenas um CEP conhecido. Outros CEPs válidos retornam "erro": true.
func servidorLocal() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != "/01001000/json/" {
			fmt.Fprintln(w, `{"erro":true}`)
			return
		}
		fmt.Fprintln(w, `{"cep":"01001-000","logradouro":"Praca da Se","localidade":"Sao Paulo","uf":"SP"}`)
	}))
}

func main() {
	online := flag.Bool("online", false, "consultar o ViaCEP pela internet")
	cep := flag.String("cep", "01001000", "CEP com oito digitos")
	flag.Parse()
	baseURL := "https://viacep.com.br/ws"
	if !*online {
		servidor := servidorLocal()
		defer servidor.Close()
		baseURL = servidor.URL
	}
	cliente := &http.Client{Timeout: 5 * time.Second}

	if !regexp.MustCompile(`^[0-9]{8}$`).MatchString(*cep) {
		log.Fatal("CEP deve conter oito digitos")
	}
	resposta, err := cliente.Get(baseURL + "/" + *cep + "/json/")
	if err != nil {
		log.Fatal(err)
	}
	defer resposta.Body.Close()
	if resposta.StatusCode != http.StatusOK {
		log.Fatalf("Status inesperado: %s", resposta.Status)
	}
	var endereco Endereco
	if err := json.NewDecoder(resposta.Body).Decode(&endereco); err != nil {
		log.Fatal(err)
	}
	if endereco.Erro {
		log.Fatal("CEP nao encontrado")
	}
	fmt.Printf("%s: %s, %s/%s\n", endereco.CEP, endereco.Logradouro, endereco.Localidade, endereco.UF)
}
