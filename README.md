# Curso de Go

Material de estudo de Go, organizado em módulos com exemplos e desafios progressivos.

## Conteúdo

### Módulo 01 — Fundação

As aulas estão organizadas na pasta [`01 - Fundação`](01%20-%20Funda%C3%A7%C3%A3o/):

1. Entendendo a primeira linha
2. Declaração e atribuição
3. Criação de tipos
4. Importando `fmt` e tipagem
5. Percorrendo arrays
6. Slices
7. Maps
8. Funções
9. Funções variáveis
10. Closures
11. Iniciando com structs
12. Composição de structs
13. Métodos em structs
14. Interfaces
15. Ponteiros
16. Quando usar ponteiros
17. Ponteiros e structs
18. Interfaces vazias
19. Type assertion
20. Generics
21. Pacotes e módulos, parte 1
22. Pacotes e módulos, parte 2
23. Pacotes e módulos, parte 3
24. Instalando pacotes
25. `for`
26. Condicionais
27. Compilando projetos

Cada aula contém um `main.go` e um `README.md` com a explicação do conceito e o comando de execução.

### Módulo 02 — Pacotes Importantes

O módulo [`02 - Pacotes Importantes`](02%20-%20Pacotes%20Importantes/README.md) segue os 22 itens do curso: arquivos, chamadas HTTP, defer, JSON, busca de CEP, servidores, headers, ServeMux, FileServer, templates, timeout, POST, requests e contextos.

As aulas 01 a 21 têm exemplos implementados, arquivos auxiliares, explicações, resultados esperados e desafios. O item 22, **Documentação para Template**, é material de consulta sem programa. A pasta `00 - Documentos` reúne os links de apoio. Consulte o índice do módulo para distinguir programas que terminam automaticamente dos servidores que ficam ativos até Ctrl+C.

Para começar, partindo da raiz do repositório:

```powershell
cd "02 - Pacotes Importantes\01 - Manipulação de arquivos"
go run main.go
```

A orientação específica está na [skill local pacotes-importantes-go](.github/skills/pacotes-importantes-go/SKILL.md), seguindo a organização da skill de Fundação.

### Módulo 03 — Context

O módulo [03 - Context](03%20-%20Context/README.md) contém cinco aulas completas: introdução aos contextos, conceitos básicos, contexto no servidor HTTP, contexto no cliente e WithValue.

Os exemplos usam apenas a biblioteca padrão e funcionam sem internet. A aula 03 mantém um servidor local; a aula 04 inicia seu próprio servidor temporário. As aulas 02 a 05 incluem testes de comportamento. Consulte a [skill context-go](.github/skills/context-go/SKILL.md) para orientação de estudo.

Para começar, partindo da raiz:

```powershell
cd "03 - Context\01 - Introdução aos contextos"
go run main.go
```

## Requisitos

- [Go](https://go.dev/dl/) 1.18 ou superior instalado (a aula 20 usa generics)
- Acesso à internet para baixar a dependência externa da aula 24 de Fundação pela primeira vez e, ao implementar as consultas externas de Pacotes Importantes, acessar os serviços usados
- Git, caso queira clonar ou atualizar este repositório

Confira a instalação com:

```powershell
go version
```

## Executar uma aula de Fundação

Nas aulas 01 a 20, 25 e 26, entre na pasta da aula desejada e execute o arquivo:

```powershell
cd "01 - Fundação\01 - Entendendo a primeira  linha"
go run main.go
```

As aulas 21 a 24 e 27 têm módulos independentes, cada um com seu próprio `go.mod`. Entre na pasta da aula e use `go run .` para incluir todos os arquivos do pacote. Por exemplo:

```powershell
cd "01 - Fundação\21 - Pacotes e módulos parte 1"
go run .
```

Os caminhos dos exemplos acima partem da raiz do repositório. Não execute os dois comandos `cd` em sequência sem voltar à raiz.

Na aula 24, execute `go mod download` antes da primeira execução. A aula 27 explica como gerar um executável com `go build`. Como as aulas são independentes, comandos de módulo devem ser executados dentro da aula correspondente, não na raiz do curso.

## Como praticar

Leia a explicação, execute o exemplo original e resolva o **Desafio prático** no README de cada aula. Use o resultado esperado para comparar a execução. Nas aulas 21 a 24 e 27 de Fundação e nas aulas de código de Pacotes Importantes, os READMEs apresentam a saída ou resposta esperada do programa original. Os desafios propõem alterações para praticar sobre os exemplos prontos.

Os materiais complementares estão em [`01 - Fundação/00 - Documentos`](01%20-%20Funda%C3%A7%C3%A3o/00%20-%20Documentos/) e no índice de apoio de [`02 - Pacotes Importantes/00 - Documentos`](02%20-%20Pacotes%20Importantes/00%20-%20Documentos/README.md).

## Criar novos módulos a partir de uma imagem

Anexe a imagem com a lista de aulas e peça:

> Crie o próximo módulo a partir desta imagem, seguindo a skill criar-modulo-go. Quero todos os exemplos implementados, READMEs e a skill específica do módulo.

A [skill criar-modulo-go](.github/skills/criar-modulo-go/SKILL.md) orienta o assistente a identificar as aulas, escolher a próxima numeração, criar exemplos completos e arquivos auxiliares, documentar, validar e atualizar os índices. O [AGENTS.md](AGENTS.md) encaminha esse tipo de pedido para a skill neste repositório.

Em um editor que suporte agentes personalizados de `.github/agents`, também é possível selecionar o [Criador de módulos Go](.github/agents/go-module-builder.agent.md) e enviar a imagem. Se ele não aparecer no seletor, use o pedido acima e indique o caminho `.github/skills/criar-modulo-go/SKILL.md`.

A criação acontece durante a conversa com o assistente. A imagem define o roteiro; os exemplos são didáticos e próprios, não transcrições dos vídeos. Itens de referência recebem documentação, sem um programa artificial. Commit e push são feitos quando solicitados.

A esteira reutiliza scripts PowerShell para criar pastas e índices e verificar estrutura, links locais, formatação e testes existentes. Os comandos estão na skill; o código dos exemplos e suas explicações continuam sendo elaborados pelo assistente. O validador mostra um resumo quando tudo passa e detalhes nas falhas.

## Repositório

O código está disponível em:

<https://github.com/luisgustavodemiranda/Curso-GO>
