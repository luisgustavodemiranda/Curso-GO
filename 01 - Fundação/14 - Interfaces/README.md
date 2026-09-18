# Item 14: interfaces

Uma interface descreve comportamentos por meio de métodos. Qualquer tipo que possua o método exigido implementa a interface implicitamente.

`Quadrado` implementa `Forma` porque possui o método `Area`.

## Executar

```powershell
go run main.go
```

## Desafio prático

Defina um tipo `Retangulo` com largura, altura e método `Area() float64`; passe-o a `imprimirArea`.

**Confira:** Com largura 4 e altura 3, deve aparecer `Area: 12`, sem alterar a interface `Forma`.
