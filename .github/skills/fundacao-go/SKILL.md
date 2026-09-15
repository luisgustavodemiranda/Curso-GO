---
name: fundacao-go
description: "Ensina, explica, depura e amplia os exercícios básicos de Go do curso 01 - Fundação. Use ao trabalhar com sintaxe, tipos, arrays, slices, maps, funções, closures, structs, interfaces, ponteiros, generics, pacotes, módulos, loops, condicionais, compilação ou erros em main.go e README.md."
argument-hint: "Informe o número da aula, o exercício, o erro ou o conceito de Go que deseja estudar."
user-invocable: true
---

# Curso Go: Fundação

## Objetivo

Acompanhar o estudante nos exercícios de `01 - Fundação`, mantendo o aprendizado progressivo e verificável. A skill deve explicar o conceito com clareza, fazer a menor alteração necessária no exercício e confirmar o resultado com ferramentas do Go.

## Quando usar

- Estudar ou revisar uma aula do curso.
- Corrigir um erro de compilação, execução ou entendimento.
- Completar um exercício existente sem apagar a intenção original.
- Criar um exemplo pequeno para um conceito já apresentado.
- Atualizar a explicação de uma aula em `README.md`.

## Procedimento

1. Identifique a aula e o objetivo do pedido.
   - Use o caminho informado pelo estudante; caso ele só forneça um número, procure a pasta correspondente em `01 - Fundação`.
   - Leia `main.go` e `README.md` quando existirem.
   - Se um arquivo esperado estiver ausente, continue com o material disponível e informe a ausência.

2. Formule a hipótese local.
   - Relacione o comportamento observado ao trecho que o controla.
   - Declare uma verificação barata que possa confirmar ou refutar a hipótese, como `go run .`, `go test ./...` ou `go build`.
   - Não faça alterações amplas antes de localizar o código responsável.

3. Explique antes de alterar.
   - Responda em português brasileiro.
   - Explique o conceito em termos adequados ao nível da aula.
   - Preserve os nomes, a organização e a intenção pedagógica existentes.
   - Ao corrigir, mostre por que o código falhava e qual regra de Go está envolvida.

4. Faça a menor alteração útil.
   - Edite somente os arquivos necessários para o pedido.
   - Prefira a sintaxe e os padrões já usados nas aulas próximas.
   - Não introduza bibliotecas, abstrações ou complexidade que o exercício ainda não ensina.
   - Para documentação, preserve os blocos de código e os comandos executáveis.

5. Valide imediatamente.
   - Formate arquivos Go com `gofmt`.
   - Para uma aula executável, rode `go run .` dentro da pasta da aula.
   - Para um projeto com módulo ou vários pacotes, use `go test ./...` e/ou `go build ./...` conforme o objetivo.
   - Compare a saída com o comportamento esperado e corrija apenas problemas relacionados ao pedido.

6. Feche o ciclo de aprendizagem.
   - Resuma o que mudou e o conceito praticado.
   - Informe a validação executada e o resultado.
   - Aponte uma pergunta curta ou um próximo experimento quando isso ajudar o estudante a consolidar o conteúdo.

## Decisões por tipo de pedido

### Dúvida conceitual
Leia o exercício relevante, explique com um exemplo mínimo e, se possível, conecte a explicação ao código existente sem modificar arquivos desnecessariamente.

### Erro de compilação ou execução
Reproduza primeiro com o comando mais estreito. Use a mensagem do compilador e o trecho responsável para corrigir a causa, depois repita exatamente a mesma verificação.

### Novo exercício ou extensão
Mantenha o escopo da aula. Se a solicitação usar um conceito ainda não apresentado, sinalize isso e ofereça uma versão compatível com a aula antes de introduzir o conceito novo.

### README ausente ou incompleto
Não invente o comportamento do programa. Derive a explicação do `main.go`, adicione apenas o contexto confirmado e inclua um comando de execução que possa ser validado.

## Critérios de conclusão

- O pedido foi atendido na aula correta.
- A explicação distingue regra de Go, comportamento do programa e escolha pedagógica.
- O código continua formatado e sem alterações não relacionadas.
- Pelo menos uma validação executável foi realizada quando o ambiente permitir.
- O resultado da validação e eventuais limitações foram comunicados ao estudante.
