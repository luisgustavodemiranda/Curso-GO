---
name: criar-modulo-go
description: "Cria módulos completos de estudo no Curso-GO a partir de imagem ou lista: exemplos implementados, documentação, validação e skill específica. Para dúvidas de aulas existentes, use a skill do módulo."
---

# Criar módulo Go

## Fluxo

1. Leia a imagem/lista: extraia título, ordem e aulas, ignorando duração e progresso. Distinga código de referências; não invente itens ilegíveis. Informe o inventário e prossiga, perguntando apenas pelo essencial que faltar. Os exemplos são próprios, não transcrições dos vídeos.
2. Confira Git, módulos existentes e a seção relevante do índice. Use o próximo número, salvo indicação do usuário. Retome destinos existentes sem sobrescrever o trabalho do estudante.
3. Use `scripts/criar-estrutura.ps1` para pastas e índices novos (comandos abaixo). A estrutura intermediária não é a entrega: implemente todos os exemplos, arquivos auxiliares e documentação na mesma tarefa, salvo pedido explícito de apenas estrutura.
4. Crie exemplos pequenos, independentes e idiomáticos. Prefira biblioteca padrão; quando necessário, use go.mod na aula, dependências versionadas e go.sum. Não entregue TODOs ou programas que só imprimem títulos. Trate erros e libere recursos.
5. Complete cada README: objetivo, funcionamento, requisitos, execução na pasta correta, resultado real e desafio com critério de conferência. Documente arquivos gerados, portas, serviços e internet. Atualize o índice do módulo e o da raiz; mantenha referências verificadas em 00 - Documentos, sem inventar ou duplicar apostilas.
6. Crie `.github/skills/nome-do-modulo-go/SKILL.md` com name, description e apenas particularidades do módulo: localização, execução, validação e decisões didáticas. Vincule no índice. Não copie regras genéricas ou de assuntos alheios.
7. Formate, valide e corrija. Execute exemplos finitos comparando saídas; para servidores, valide handlers ou uma chamada local e encerre processos iniciados. Use testes de comportamento quando úteis. Confira manualmente a correspondência imagem/índice e os metadados da skill.

## Scripts reutilizáveis

Execute na raiz do curso. Exemplo ilustrativo; substitua nome e aulas pela entrada real:

```powershell
& .github/skills/criar-modulo-go/scripts/criar-estrutura.ps1 -Nome 'Novo assunto' -Aulas @('Introdução', 'Exemplo', 'Referências') -Referencias @(3)
```

O script escolhe o próximo número, normaliza nomes Windows, gera pastas e bases de README e recusa módulo já existente. Não gera código, skill ou altera o índice da raiz. `-Raiz` permite uma pasta de teste. `-Referencias` indica posições documentais; omita se todas forem código. Conclua os textos em preparação antes da entrega.

Depois de implementar e formatar:

```powershell
& .github/skills/criar-modulo-go/scripts/validar-modulo.ps1 -Modulo 'NN - Novo assunto' -Referencias @(3) -Testes
```

Confere READMEs, main.go das aulas de código, links Markdown locais simples e gofmt. `-Testes` executa testes existentes por aula com timeout de 20 segundos; a compilação/download não está abrangida por esse timeout. Mostra resumo em sucesso e detalhes nas falhas. Não executa go run nem servidores, não verifica links externos, resultados, semântica ou skills. Um resultado OK não substitui essas verificações.

Se a política local bloquear arquivos ps1, use uma exceção somente no processo, sem mudar a política permanente. Para validar:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File .github/skills/criar-modulo-go/scripts/validar-modulo.ps1 -Modulo '03 - Context' -Testes
```

Para passar listas ao criador, execute o comando com arrays em uma sessão temporária iniciada com `powershell -NoProfile -ExecutionPolicy Bypass` e encerre-a com `exit`. Depois de alterar os scripts, rode `scripts/testar-scripts.ps1`: ele verifica criação e falhas esperadas em uma pasta temporária, sem alterar aulas reais. Não repita esses testes a cada módulo se os scripts não mudaram.

## Decisões de execução

- Use `go run main.go` para arquivo independente; com vários arquivos, módulo de aula e `go run .` ou lista explícita de arquivos. Não crie módulo Go na raiz do curso por conveniência.
- HTTP: servidores locais, clientes com timeout, simulação identificada e integração real quando essencial ao assunto. Bancos e serviços: configuração reproduzível sem credenciais reais. Declare o que não pôde executar; não substitua silenciosamente o tema por simulação.
- Contexto econômico: não releia arquivos já conhecidos e inalterados. Consulte trechos e diffs pertinentes; abra referências só quando necessárias. Use os scripts sem reimprimir seu código. Após passar as verificações, repita somente se houver mudança ou falha que justifique.
- Preserve explicações didáticas e testes úteis. A economia deve vir de repetição mecânica, não de entregar material incompleto.

## Entrega e commit

Informe caminhos, quantidade de aulas de código/referência, validações, limitações e um comando inicial. Revise o diff. Se o commit desta entrega já estiver autorizado, inclua apenas seus arquivos relacionados e informe hash e estado final. Caso contrário, termine perguntando **“Quer que eu faça o commit dessas alterações agora?”** Não faça push sem autorização específica.
