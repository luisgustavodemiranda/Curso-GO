# Item 09: funções variáveis

Uma função variádica recebe zero ou mais valores usando `...int`. Dentro da função, o parâmetro se comporta como um slice.

O exemplo soma qualquer quantidade de números.

## Executar

```powershell
go run main.go
```

## Desafio prático

Experimente `somar()` e depois `somar(10, 20, 30)`.

**Confira:** Os resultados devem ser 0 e 60, respectivamente.
