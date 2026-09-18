---
name: pacotes-importantes-go
description: "Orienta o estudo, a implementação e a revisão das aulas de 02 - Pacotes Importantes no Curso-GO, envolvendo arquivos, JSON, clientes e servidores HTTP, CEP, templates e contextos. Use para exercícios e documentação desse módulo."
---

# Curso Go: Pacotes Importantes

## Contexto do projeto

- O módulo está em `02 - Pacotes Importantes`; consulte seu `README.md` para localizar a aula pelo número e título.
- As aulas 01 a 21 têm exemplos implementados; preserve seu comportamento ao explicar ou ampliar um exercício. As aulas 11 e 14 a 16 incluem arquivos auxiliares em public ou templates.
- O item 22 é documentação de templates, sem programa. `00 - Documentos` aponta para a apostila já existente em Fundação.
- As aulas são independentes e usam `go run main.go`, sem módulo Go compartilhado. Nas aulas que têm testes, use `go test main.go main_test.go`. Verifique os arquivos atuais antes de escolher comandos.
- As aulas 06, 07, 09, 10, 11 e 15 são servidores persistentes em localhost:8080. As aulas 05, 08 e 09 simulam CEP por padrão; a flag -online consulta o ViaCEP. Preserve a possibilidade de estudar sem internet.
- Não confunda a numeração dessas aulas com a do módulo Fundação. Se o pedido só mencionar um número, use o contexto do arquivo aberto; pergunte qual módulo apenas se continuar ambíguo.

## Acompanhar o estudo

Responda em português brasileiro. Leia o README e o código da aula solicitada e, quando necessário, a aula anterior. Em dúvidas, explique com o exemplo local e ofereça uma pista ou experimento. Quando o estudante pedir implementação, complete a aula solicitada com código pequeno e executável, preservando a sequência do curso.

Ao alterar um exemplo, atualize a explicação e o resultado esperado no README. Se adicionar uma aula ainda incompleta, identifique esse estado no índice; não a marque como concluída apenas porque compila.

Mantenha objetivo, comando de execução, resultado esperado e desafio verificável. Documente arquivos gerados, diretório de execução, rotas, porta e dependência de internet conforme o exemplo. Se a estrutura ou o índice mudar, sincronize o README do módulo e o da raiz.

## Decisões por assunto

- **Arquivos e defer:** trabalhe com arquivos de exercício na pasta da aula ou em diretório temporário. Trate falhas de abertura antes de registrar o fechamento; explique a ordem de execução dos defers.
- **Clientes HTTP e CEP:** trate erros de transporte, status HTTP, decodificação e resposta de CEP inexistente separadamente. Feche o corpo após receber uma resposta válida. Use timeout finito nas chamadas externas; nas aulas 18 e 21, explique timeout e cancelamento explicitamente.
- **Servidores:** use endereço local e documente como iniciar, consultar e encerrar. Defina headers e status antes do corpo. Para FileServer, sirva uma pasta de conteúdo estático, não o repositório.
- **JSON:** use campos exportados e tags coerentes com o exemplo; trate erros de codificação e decodificação.
- **Templates:** use `html/template` para HTML e `text/template` para texto. Registre FuncMap antes de Parse, confira nomes ao compor templates e trate erros de execução. Reserve Must para a inicialização de templates conhecidos, explicando seu panic.
- **POST e requests customizadas:** prefira um servidor local para observar método, headers e corpo sem gravar dados em serviços externos.

## Validação proporcional

Formate o código alterado com gofmt. Execute exemplos finitos e compare a saída; para servidores, valide uma rota com requisição local ou use net/http/httptest quando um teste for útil, encerrando o processo iniciado após a verificação. Use servidor local com atraso para timeout e cancelamento.

Se houver go.mod na aula, execute os comandos de módulo dentro dela. Sem go.mod, use o arquivo explícito. Diferencie falha de rede ou porta ocupada de defeito do código. Prefira a biblioteca padrão; adicione dependência externa apenas quando o exercício precisar dela.

Ao concluir, informe o conceito praticado, os arquivos alterados, a validação realizada e o que ainda ficou pendente.
