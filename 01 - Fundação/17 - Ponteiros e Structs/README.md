# Item 17: ponteiros e structs

Este item combina structs e ponteiros. Um ponteiro para uma struct permite alterar seus campos diretamente usando a notação `config.Nome`.

O exemplo define uma `Pessoa` e passa seu endereço para a função `atualizarNome`. Como a função recebe `*Pessoa`, ela altera o campo `Nome` da struct original.

## Executar

```powershell
go run main.go
```
