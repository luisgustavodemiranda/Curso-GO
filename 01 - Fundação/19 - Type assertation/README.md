# Item 19: type assertion

Uma type assertion verifica se um valor armazenado em uma interface possui um tipo específico.

A forma `nome, ok := valor.(string)` evita pânico: `ok` informa se o valor dinâmico da interface é uma `string`. Se for, `nome` recebe essa string; caso contrário, recebe o valor zero `""` e `ok` é `false`. A assertion não converte números ou outros tipos em texto.

## Executar

```powershell
go run main.go
```

## Desafio prático

Acrescente `descrever(true)` e `descrever("")`.

**Confira:** As novas linhas devem ser `Outro tipo: true` e `Texto: `; a string vazia ainda satisfaz a assertion.
