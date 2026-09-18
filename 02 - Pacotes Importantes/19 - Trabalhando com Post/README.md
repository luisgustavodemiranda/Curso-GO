# Aula 19: Trabalhando com Post

## Objetivo

Enviar dados no corpo de uma requisição POST.

## Como funciona

Um servidor temporário recebe POST, decodifica a struct e responde JSON com status 201. O cliente codifica a pessoa e envia com `Client.Post`. Tudo ocorre localmente, sem persistir dados.

## Executar

Dentro da pasta desta aula:

```powershell
go run main.go
```

## Resultado esperado

```text
Status: 201 Created
{"nome":"Ana"}
```

## Desafio prático

Adicione idade ao JSON enviado e à struct recebida; confira a resposta.

[Voltar ao índice do módulo](../README.md)
