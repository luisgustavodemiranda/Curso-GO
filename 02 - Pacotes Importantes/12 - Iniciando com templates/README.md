# Aula 12: Iniciando com templates

## Objetivo

Gerar texto a partir de dados usando text/template.

## Como funciona

`template.New` nomeia o template, `Parse` analisa seu texto e `Execute` recebe o destino e os dados. `{{.Nome}}` acessa um campo exportado da struct.

## Executar

Dentro da pasta desta aula:

```powershell
go run main.go
```

## Resultado esperado

```text
Ola, Ana!
```

## Desafio prático

Acrescente uma idade à struct e ao texto do template.

[Voltar ao índice do módulo](../README.md)
