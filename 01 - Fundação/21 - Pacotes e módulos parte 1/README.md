# Item 21: pacotes e módulos, parte 1

Um pacote pode reunir vários arquivos Go na mesma pasta. Aqui, `main.go` chama `mensagem`, definida em `mensagem.go`; ambos declaram `package main`.

O arquivo `go.mod` identifica o módulo como `example.com/curso-go/aula21`. Esse é um nome didático: não é necessário publicar o projeto nesse endereço.

## Executar

Dentro desta pasta:

```powershell
go run .
```

O ponto inclui os arquivos do pacote. `go run main.go` sozinho falha neste exemplo, porque não inclui `mensagem.go`.

## Criar um módulo do zero

Em uma pasta nova, ainda sem `go.mod`, use `go mod init example.com/curso-go/aula21`. Nesta aula o arquivo já está pronto; não é necessário repetir a inicialização.

## Saída esperada

```text
Ola, Ana!
```

## Desafio prático

Crie `despedida.go` no mesmo pacote, com uma função `despedida(nome string) string`, e chame-a no `main`.

**Confira:** a nova linha deve mostrar `Ate logo, Ana!`.
