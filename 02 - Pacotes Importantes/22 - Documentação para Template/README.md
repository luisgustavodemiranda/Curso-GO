# Item 22: documentação para Template

Material de consulta para as aulas 12 a 17. Este item não contém `main.go` porque é uma referência, conforme a lista do curso.

## Referências oficiais

- [text/template](https://pkg.go.dev/text/template): sintaxe, ações, pipelines, `if`, `range`, `with` e composição.
- [html/template](https://pkg.go.dev/html/template): renderização de HTML com escape contextual dos dados.

## Roteiro de consulta

1. Localize os exemplos de `Parse` e `Execute`.
2. Compare `ParseFiles` e `ExecuteTemplate`.
3. Consulte `FuncMap` para registrar funções antes de analisar o template.
4. Releia as ações `define` e `template` ao estudar composição.

## Desafio de leitura

Encontre um exemplo de `range` e adapte-o na aula 12 para listar três nomes. Confira se o resultado contém os três valores na ordem fornecida.

[Voltar ao índice do módulo](../README.md)
