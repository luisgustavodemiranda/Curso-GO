# Curso de Go

Material de estudo da formação **Fundação Go**, com exemplos progressivos da linguagem Go.

## Conteúdo

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

## Requisitos

- [Go](https://go.dev/dl/) 1.18 ou superior instalado (a aula 20 usa generics)
- Acesso à internet para baixar a dependência externa da aula 24 pela primeira vez
- Git, caso queira clonar ou atualizar este repositório

Confira a instalação com:

```powershell
go version
```

## Executar uma aula

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

Leia a explicação, execute o exemplo original e resolva o **Desafio prático** no README de cada aula. Use o critério **Confira** para comparar o resultado. Nas aulas 21 a 24 e 27, os READMEs também apresentam a saída esperada do programa original.

Os materiais complementares estão em [`01 - Fundação/00 - Documentos`](01%20-%20Funda%C3%A7%C3%A3o/00%20-%20Documentos/).

## Repositório

O código está disponível em:

<https://github.com/luisgustavodemiranda/Curso-GO>
