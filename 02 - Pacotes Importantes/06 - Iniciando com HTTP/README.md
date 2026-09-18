# Aula 06: Iniciando com HTTP

## Objetivo

Criar um servidor HTTP com net/http.

## Como funciona

Um handler recebe `http.ResponseWriter` e `*http.Request`. Ele responde na rota `/` e devolve 404 nos demais caminhos. O servidor fica ativo até Ctrl+C.

## Executar

Dentro da pasta desta aula:

```powershell
go run main.go
```

O processo permanece ativo em `localhost:8080`. Em outro terminal, consulte:

```powershell
curl.exe -i "http://localhost:8080/"
```

Execute uma aula de servidor por vez e encerre com **Ctrl+C**. Se a porta estiver ocupada, encerre o outro servidor ou altere `Addr` no código.

## Resultado esperado

Resultado da consulta HTTP (não da inicialização do servidor):

```text
Ola, HTTP!
```

## Desafio prático

Adicione uma rota /sobre e compare com uma rota inexistente.

[Voltar ao índice do módulo](../README.md)
