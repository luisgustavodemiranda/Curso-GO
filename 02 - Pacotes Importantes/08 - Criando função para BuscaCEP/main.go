package main

import (
	"encoding/json"
	"errors"
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

var (
	errCEPInvalido    = errors.New("CEP deve conter oito digitos")
	errCEPInexistente = errors.New("CEP nao encontrado")
	formatoCEP        = regexp.MustCompile(`^[0-9]{8}$`)
)

// Receber cliente e URL permite testar sem depender de um serviço externo.
func buscarCEP(cliente *http.Client, baseURL, cep string) (Endereco, error) {
	if !formatoCEP.MatchString(cep) {
		return Endereco{}, errCEPInvalido
	}
	resposta, err := cliente.Get(baseURL + "/" + cep + "/json/")
	if err != nil {
		return Endereco{}, fmt.Errorf("consultar CEP: %w", err)
	}
	defer resposta.Body.Close()
	if resposta.StatusCode != http.StatusOK {
		return Endereco{}, fmt.Errorf("servico de CEP retornou %s", resposta.Status)
	}
	var endereco Endereco
	if err := json.NewDecoder(resposta.Body).Decode(&endereco); err != nil {
		return Endereco{}, fmt.Errorf("decodificar CEP: %w", err)
	}
	if endereco.Erro {
		return Endereco{}, errCEPInexistente
	}
	return endereco, nil
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

	endereco, err := buscarCEP(cliente, baseURL, *cep)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s: %s, %s/%s\n", endereco.CEP, endereco.Logradouro, endereco.Localidade, endereco.UF)
}
