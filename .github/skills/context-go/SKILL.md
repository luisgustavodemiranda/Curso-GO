---
name: context-go
description: "Orienta o estudo, a depuração e a extensão das cinco aulas de 03 - Context no Curso-GO: cancelamento, prazos, contexto de servidor e cliente HTTP e WithValue."
---

# Estudar Context no Curso-GO

## Organização

Consulte `03 - Context/README.md` para os caminhos das cinco aulas. Cada uma é independente, com main.go e README, sem go.mod. Execute `go run main.go` dentro da aula. As aulas 02 a 05 têm testes executáveis com `go test -timeout 20s main.go main_test.go`.

A aula 03 mantém um servidor em localhost:8080 até Ctrl+C. A aula 04 usa um servidor temporário próprio; não precisa que a aula 03 esteja rodando. Os demais exemplos não usam rede. Todos funcionam sem internet.

## Orientação didática

Leia o código e o README da aula solicitada antes de explicar ou editar. Responda em português e preserve a progressão: introdução, conceitos básicos, servidor, cliente e valores. Não refatore outras aulas para compartilhar funções; a independência é intencional.

Explique que cancelar um contexto sinaliza a interrupção: a operação precisa observar Done ou usar uma API que respeite o contexto. Não descreva cancelamento como encerramento forçado de goroutines. Diferencie Canceled, DeadlineExceeded e o timeout do cliente HTTP.

Passe ctx como primeiro parâmetro e preserve o contexto recebido ao derivar filhos. Em handlers, derive de r.Context, não de Background. Libere os recursos dos contextos derivados com cancel. Mostre que cancelar um filho não cancela o pai.

Na aula de servidor, diferencie expiração do prazo local (resposta 504 escolhida pelo exemplo) de cancelamento da requisição do cliente (abandono do trabalho). O pacote context não escolhe códigos HTTP.

Na aula WithValue, mantenha chaves de tipo privado e leitura segura. Use valores para metadados da requisição; não transforme o contexto em armazenamento de configuração ou parâmetros opcionais.

## Alterar e validar

Atualize explicação, comandos, saídas e desafios quando mudar o exemplo. Sincronize o índice se houver nova aula ou mudança de execução. Formate com gofmt, execute exemplos finitos e rode os testes afetados. Para servidor, teste o handler ou faça uma requisição local e encerre qualquer processo iniciado.

Prefira canais de sincronização e contextos já cancelados ou expirados para testar propagação; evite testes baseados em tempos exatos. Use limite de execução para detectar bloqueios. Informe validações e limitações sem confundir falha do ambiente com defeito do código.
