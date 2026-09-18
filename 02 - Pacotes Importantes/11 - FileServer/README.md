# Aula 11: FileServer

## Objetivo

Servir arquivos estáticos com http.FileServer.

## Como funciona

`http.FileServer` serve apenas a pasta `public`, que contém `index.html`. Os caminhos são relativos ao diretório de execução: execute dentro da aula.

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
Ao abrir / no navegador: FileServer em Go
```

## Desafio prático

Crie public/sobre.html e acesse /sobre.html.

[Voltar ao índice do módulo](../README.md)
