# Item 15: ponteiros

Um ponteiro guarda o endereço de uma variável. `&numero` obtém o endereço e `*valor` acessa o conteúdo apontado.

O exemplo usa um ponteiro para alterar o número dentro da função.

## Executar

```powershell
go run main.go
```

## Desafio prático

Chame `dobrar(&numero)` duas vezes antes de imprimir o número.

**Confira:** O valor inicial 10 deve se tornar 40.
