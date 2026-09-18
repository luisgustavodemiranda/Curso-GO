# Item 24: instalando pacotes

Esta aula usa uma dependência externa real: [github.com/google/uuid](https://github.com/google/uuid). O exemplo interpreta um UUID fixo e mostra sua versão, mantendo a saída reproduzível.

Diferentemente de `fmt`, esse pacote não acompanha a biblioteca padrão do Go.

## Executar

Dentro desta pasta, com acesso à internet no primeiro download:

```powershell
go mod download
go run .
```

O `go.mod` já declara a versão `v1.6.0`. O `go.sum` registra checksums usados para verificar os módulos baixados; os dois arquivos devem acompanhar o código no Git.

## Praticar a instalação

O comando usado para adicionar essa dependência a um módulo é:

```powershell
go get github.com/google/uuid@v1.6.0
go mod tidy
go list -m all
```

Você pode executá-lo nesta pasta: como a versão já está declarada, não é preciso apagar os arquivos para praticar. `go get` adiciona ou altera dependências; o import permite usar o pacote no código. `go mod tidy` ajusta as dependências aos imports existentes.

## Saída esperada

```text
UUID: 550e8400-e29b-41d4-a716-446655440000
Versao: VERSION_4
```

## Desafio prático

Substitua o UUID por `"invalido"` e observe o tratamento do erro de `uuid.Parse`.

**Confira:** o programa deve imprimir uma linha iniciada por `Erro:`, sem imprimir UUID e versão. Depois restaure o valor original.
