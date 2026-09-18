# Item 22: pacotes e módulos, parte 2

Esta aula separa a saudação em um pacote próprio. A pasta `saudacao` contém `package saudacao`; o programa continua em `package main`.

O import `example.com/curso-go/aula22/saudacao` combina o nome do módulo, declarado em `go.mod`, com a subpasta. O Go encontra esse pacote no próprio projeto.

`Mensagem` começa com maiúscula para poder ser usada por outro pacote. Não é preciso instalar esse pacote local com `go get`.

## Executar

Dentro desta pasta:

```powershell
go run .
go list ./...
```

O segundo comando lista os pacotes `example.com/curso-go/aula22` e `example.com/curso-go/aula22/saudacao`.

## Saída esperada do programa

```text
Ola, Ana!
```

## Desafio prático

Adicione `Despedida` ao pacote `saudacao` e chame `saudacao.Despedida("Ana")` no programa.

**Confira:** a nova linha deve mostrar `Ate logo, Ana!`. Experimente iniciar o nome da função com minúscula e observe o erro ao tentar acessá-la pelo `main`; depois restaure a maiúscula.
