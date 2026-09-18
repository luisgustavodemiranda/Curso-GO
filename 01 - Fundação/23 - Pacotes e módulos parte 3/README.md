# Item 23: pacotes e módulos, parte 3

O pacote próprio `saudacao` agora usa `strings`, da biblioteca padrão. A função exportada `Mensagem` chama a auxiliar `normalizar`, que remove espaços nas extremidades e transforma o nome em maiúsculas.

Como `normalizar` começa com minúscula, seu uso fica restrito ao pacote `saudacao`. Isso permite separar a interface pública dos detalhes internos.

Cada arquivo declara os imports que utiliza. O pacote `strings` acompanha o Go e não precisa ser adicionado como dependência externa.

## Executar e inspecionar

Dentro desta pasta:

```powershell
go run .
go mod tidy
go list -m
```

`go mod tidy` ajusta as dependências conforme os imports. Como esta aula só usa um pacote local e a biblioteca padrão, não há dependências externas nem necessidade de `go.sum`. O último comando mostra `example.com/curso-go/aula23`.

## Saída esperada

```text
Ola, ANA!
```

## Desafio prático

Passe `"  Maria  "` para `saudacao.Mensagem` no `main`.

**Confira:** a saída deve ser `Ola, MARIA!`, sem os espaços extras.
