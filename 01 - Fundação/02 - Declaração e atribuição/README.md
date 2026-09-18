# Item 02: declaração e atribuição

O exemplo apresenta três formas de criar variáveis:

- `var nome string` declara o tipo antes do valor.
- `var idade int = 20` declara a variável com valor inicial.
- `cidade := "Sao Paulo"` usa a declaração curta e deixa o Go inferir o tipo.

A atribuição `idade = 21` altera uma variável que já existe.

## Executar

```powershell
go run main.go
```

## Desafio prático

Altere a atribuição final de `idade` para 22 e a cidade para Recife.

**Confira:** A saída deve mostrar Maria, 22 e Recife, em linhas separadas.
