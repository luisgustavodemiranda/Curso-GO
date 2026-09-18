# Aula 20: Customizando Request

## Objetivo

Montar uma requisição com http.NewRequest e enviá-la com Client.Do.

## Como funciona

`http.NewRequest` cria um PUT com corpo JSON e cabeçalhos. `Client.Do` envia a requisição. O servidor temporário devolve os dados recebidos para você conferir a personalização.

## Executar

Dentro da pasta desta aula:

```powershell
go run main.go
```

## Resultado esperado

```text
Metodo: PUT
X-Curso: Go
Content-Type: application/json
Corpo: {"nome":"Maria"}
```

## Desafio prático

Troque PUT por PATCH e X-Curso por outro valor; confira os dados devolvidos.

[Voltar ao índice do módulo](../README.md)
