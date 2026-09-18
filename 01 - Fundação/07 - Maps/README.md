# Item 07: maps

Maps armazenam pares de chave e valor. Neste exemplo, nomes são associados a idades.

A leitura `idade, existe := idades["Ana"]` também informa se a chave foi encontrada.

## Executar

```powershell
go run main.go
```

## Desafio prático

Consulte uma chave ausente usando `idade, existe := idades["Pedro"]` em vez da consulta de Ana. Ajuste também o rótulo impresso.

**Confira:** O valor deve ser 0 e `existe` deve ser `false`.
