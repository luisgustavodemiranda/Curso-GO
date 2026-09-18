# Aula 10: ServeMux

## Objetivo

Organizar rotas usando um http.ServeMux próprio.

## Como funciona

`http.NewServeMux` mantém duas rotas. O handler `/` verifica o caminho para não transformar rotas desconhecidas em uma página de sucesso. `/sobre` mostra o assunto do curso.

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
/ → Pagina inicial
/sobre → Curso de Go — Pacotes Importantes
/inexistente → HTTP 404
```

## Desafio prático

Adicione /contato sem alterar o resultado de /inexistente.

[Voltar ao índice do módulo](../README.md)
