# Aula 03: Defer

## Objetivo

Entender a execução adiada de funções com defer.

## Como funciona

Os `defer` são registrados em `exemplo` e executados quando essa função retorna, em ordem inversa. Eles não esperam necessariamente o encerramento do programa.

## Executar

Dentro da pasta desta aula:

```powershell
go run main.go
```

## Resultado esperado

```text
inicio
fim
segundo defer
primeiro defer
```

## Desafio prático

Adicione um terceiro defer e preveja a ordem antes de executar.

[Voltar ao índice do módulo](../README.md)
