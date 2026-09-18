# Item 27: compilando projetos

`go run .` compila e executa o programa durante o desenvolvimento. `go build` gera um executável que pode ser iniciado diretamente, sem chamar o Go novamente.

Esta aula tem seu próprio `go.mod`. Execute os comandos dentro desta pasta.

## Executar durante o estudo

```powershell
go run .
```

## Compilar e executar no Windows

```powershell
New-Item -ItemType Directory -Force bin | Out-Null
go build -o bin/curso-go.exe .
.\bin\curso-go.exe
```

O executável fica na pasta `bin`, ignorada pelo Git. Ele é gerado para o sistema e a arquitetura usados na compilação; esse arquivo Windows não é um executável Linux.

No Linux ou macOS, use:

```sh
mkdir -p bin
go build -o bin/curso-go .
./bin/curso-go
```

## Saída esperada

```text
Projeto compilado com sucesso!
```

## Desafio prático

Após compilar, altere a mensagem no código e execute o binário antigo novamente. Em seguida, repita `go build` e execute o binário atualizado.

**Confira:** a mensagem só muda no executável depois de uma nova compilação.
