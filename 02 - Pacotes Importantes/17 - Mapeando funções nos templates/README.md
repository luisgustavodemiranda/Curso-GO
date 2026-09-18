# Aula 17: Mapeando funções nos templates

## Objetivo

Registrar funções auxiliares com template.FuncMap.

## Como funciona

`template.FuncMap` associa o nome `maiusculas` a `strings.ToUpper`. O mapa é registrado antes de Parse para que o analisador reconheça a função.

## Executar

Dentro da pasta desta aula:

```powershell
go run main.go
```

## Resultado esperado

```text
Ola, ANA!
```

## Desafio prático

Registre uma função minusculas e compare com maiusculas.

[Voltar ao índice do módulo](../README.md)
