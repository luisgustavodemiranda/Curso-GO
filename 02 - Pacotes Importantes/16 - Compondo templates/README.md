# Aula 16: Compondo templates

## Objetivo

Reutilizar partes da página com define e template.

## Como funciona

`ParseGlob` carrega os três arquivos da pasta templates. `ExecuteTemplate` escolhe `pagina`, que usa os blocos `cabecalho` e `rodape`. A saída é HTML no terminal; esta aula não inicia servidor.

## Executar

Dentro da pasta desta aula:

```powershell
go run main.go
```

## Resultado esperado

```text
O HTML contém Curso de Go, Ola, Ana! e Estudando composicao de templates.
```

## Desafio prático

Adicione um bloco menu e inclua-o na página com template.

[Voltar ao índice do módulo](../README.md)
