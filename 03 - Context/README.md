# Módulo 03 — Context

Cinco aulas com exemplos completos e independentes, criados a partir dos temas da imagem do curso. Os exemplos são próprios, não transcrições dos vídeos.

## Aulas

1. [Introdução aos contextos](01%20-%20Introdu%C3%A7%C3%A3o%20aos%20contextos/README.md)
2. [Entendendo conceitos básicos](02%20-%20Entendendo%20conceitos%20b%C3%A1sicos/README.md)
3. [Context utilizando server HTTP](03%20-%20Context%20utilizando%20server%20HTTP/README.md)
4. [Context no lado do Client](04%20-%20Context%20no%20lado%20do%20Client/README.md)
5. [Context WithValue](05%20-%20Context%20WithValue/README.md)

## Começar

Partindo da raiz do repositório:

```powershell
cd "03 - Context\01 - Introdução aos contextos"
go run main.go
```

Cada aula usa apenas a biblioteca padrão, em um arquivo main.go, sem go.mod. Execute os comandos dentro da aula. Mantém-se o requisito do curso de Go 1.18 ou superior.

As aulas 01, 02, 04 e 05 terminam sozinhas. A aula 03 abre um servidor em localhost:8080 até Ctrl+C. A aula 04 cria e encerra seu próprio servidor em uma porta local livre. Nenhuma aula precisa de internet nem grava arquivos.

As aulas 02 a 05 incluem testes: use `go test -timeout 20s main.go main_test.go` em cada pasta. Não é necessário iniciar os servidores antes dos testes.

## Apoio

- [Referências oficiais](00%20-%20Documentos/README.md).
- [Skill context-go](../.github/skills/context-go/SKILL.md).
- [Índice do curso](../README.md).
