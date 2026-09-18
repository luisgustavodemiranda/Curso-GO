# Item 12: composição de structs

A composição permite incluir uma struct dentro de outra. `Pessoa` incorpora `Endereco` e pode acessar diretamente seus campos.

Essa abordagem favorece a reutilização sem criar uma hierarquia de classes.

## Executar

```powershell
go run main.go
```

## Desafio prático

Acesse a cidade tanto por `pessoa.Cidade` quanto por `pessoa.Endereco.Cidade`.

**Confira:** As duas expressões devem produzir `Recife`.
