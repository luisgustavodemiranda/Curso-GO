# Aula 09: Finalizando resposta para o servidor HTTP

## Objetivo

Integrar a busca de CEP a um handler que responde JSON.

## Como funciona

O handler GET `/cep` usa `buscarCEP` e retorna JSON: 200 em sucesso, 400 para formato inválido ou parâmetro ausente, 404 para CEP inexistente, 405 para outro método e 502 para falha da API consultada. Por padrão, a API é simulada localmente.

## Executar

Dentro da pasta desta aula:

```powershell
go run main.go
```

O processo permanece ativo em `localhost:8080`. Em outro terminal, consulte:

```powershell
curl.exe -i "http://localhost:8080/cep?cep=01001000"
```

Execute uma aula de servidor por vez e encerre com **Ctrl+C**. Se a porta estiver ocupada, encerre o outro servidor ou altere `Addr` no código.

## Consulta real de CEP

O modo padrão funciona sem internet e reconhece apenas o CEP `01001000`. Para consultar o [ViaCEP](https://viacep.com.br/) real:

```powershell
go run main.go -online
```

A chamada externa tem timeout de cinco segundos. Os dados reais podem conter acentos e diferir da amostra local. As aulas 05 e 08 encerram com erro para CEP inválido/inexistente; a aula 09 comunica o erro pelo status HTTP.

## Resultado esperado

Resultado da consulta HTTP (não da inicialização do servidor):

```text
{"cep":"01001-000","logradouro":"Praca da Se","localidade":"Sao Paulo","uf":"SP"}
```

## Desafio prático

Consulte a rota sem cep, com um CEP inválido e com um inexistente; confira status e corpo.

[Voltar ao índice do módulo](../README.md)
