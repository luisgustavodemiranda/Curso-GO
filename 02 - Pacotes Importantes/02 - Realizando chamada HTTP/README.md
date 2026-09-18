# Aula 02: Realizando chamada HTTP

## Objetivo

Fazer uma requisição GET com net/http e ler a resposta.

## Como funciona

`httptest.NewServer` inicia um servidor de demonstração em uma porta local livre. O cliente faz GET, confere o status e lê o corpo. `defer` fecha a resposta e o servidor ao terminar. Não exige internet.

## Executar

Dentro da pasta desta aula:

```powershell
go run main.go
```

## Resultado esperado

```text
Status: 200 OK
Ola do servidor HTTP!
```

## Desafio prático

Mude a resposta local para HTTP 404 e observe o tratamento do status; depois restaure 200.

[Voltar ao índice do módulo](../README.md)
