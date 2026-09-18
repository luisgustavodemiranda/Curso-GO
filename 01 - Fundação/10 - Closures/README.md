# Item 10: closures

Uma closure é uma função que mantém acesso às variáveis do escopo onde foi criada.

`criarContador` devolve uma função que continua acessando e atualizando `contador` entre as chamadas.

## Executar

```powershell
go run main.go
```

## Desafio prático

Crie dois contadores com duas chamadas a `criarContador`. Chame o primeiro duas vezes e o segundo uma vez.

**Confira:** As chamadas devem retornar 1, 2 e 1: cada contador mantém seu próprio estado.
