# Aula 01: Manipulação de arquivos

## Objetivo

Criar, escrever e ler arquivos usando os e io.

## Como funciona

`os.Create` abre um arquivo temporário para leitura e escrita. Depois de `WriteString`, `Seek` reposiciona o cursor e `io.ReadAll` lê o conteúdo. A função auxiliar permite executar os `defer` antes de o `main` tratar o erro. O arquivo e a pasta temporária são removidos ao terminar.

## Executar

Dentro da pasta desta aula:

```powershell
go run main.go
```

## Resultado esperado

```text
Estudando arquivos em Go!
```

## Desafio prático

Grave duas linhas e confira se ambas são lidas após o Seek.

[Voltar ao índice do módulo](../README.md)
