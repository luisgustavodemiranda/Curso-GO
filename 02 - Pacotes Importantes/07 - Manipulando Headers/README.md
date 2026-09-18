# Aula 07: Manipulando Headers

## Objetivo

Ler cabeçalhos da requisição e definir cabeçalhos da resposta.

## Como funciona

O handler lê `User-Agent` e define `Content-Type` e `X-Curso` antes de escrever o status e o corpo. O valor de User-Agent depende do cliente que faz a chamada.

## Executar

Dentro da pasta desta aula:

```powershell
go run main.go
```

O processo permanece ativo em `localhost:8080`. Em outro terminal, consulte:

```powershell
curl.exe -i -A "EstudanteGo" "http://localhost:8080/"
```

Execute uma aula de servidor por vez e encerre com **Ctrl+C**. Se a porta estiver ocupada, encerre o outro servidor ou altere `Addr` no código.

## Resultado esperado

Resultado da consulta HTTP (não da inicialização do servidor):

```text
User-Agent: EstudanteGo
```

## Desafio prático

Adicione o header X-Aula com valor 07 e confira com curl.exe -i.

[Voltar ao índice do módulo](../README.md)
