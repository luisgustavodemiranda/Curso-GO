# Aula 21: Trabalhando com HTTP usando Contextos

## Objetivo

Cancelar requisições HTTP com context.

## Como funciona

O contexto expira em 100 ms, enquanto o servidor levaria um segundo para responder. `NewRequestWithContext` vincula a requisição ao contexto. O cliente reconhece `context.DeadlineExceeded`, e o handler observa o cancelamento pelo contexto da requisição.

## Executar

Dentro da pasta desta aula:

```powershell
go run main.go
```

## Resultado esperado

```text
Requisicao cancelada: prazo do contexto esgotado
```

## Desafio prático

Reduza o atraso do servidor abaixo de 100 ms e adapte o ramo de sucesso para ler e imprimir a resposta.

[Voltar ao índice do módulo](../README.md)
