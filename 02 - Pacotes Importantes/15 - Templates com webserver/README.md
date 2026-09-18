# Aula 15: Templates com webserver

## Objetivo

Renderizar HTML dentro de um handler HTTP.

## Como funciona

O servidor carrega `templates/pagina.html` com `html/template`. O parâmetro `nome` da URL preenche a página; sem parâmetro, usa Ana. Um buffer evita enviar uma resposta parcial antes de verificar o erro de renderização.

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
No navegador: Ola, Ana!
```

## Desafio prático

Use ?nome=%3Cb%3EAna%3C%2Fb%3E e confirme que as tags aparecem como texto.

[Voltar ao índice do módulo](../README.md)
