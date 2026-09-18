# Item 18: interfaces vazias

A interface vazia, escrita como `interface{}`, não exige nenhum método. Por isso, pode receber valores de diferentes tipos.

O verbo `%v` permite imprimir o valor sem conhecer seu tipo antecipadamente.

## Executar

```powershell
go run main.go
```

## Desafio prático

Altere o formato para `"Valor: %v, tipo: %T\n"` e passe `valor` para os dois verbos.

**Confira:** Os tipos das três chamadas devem aparecer como `string`, `int` e `bool`.
