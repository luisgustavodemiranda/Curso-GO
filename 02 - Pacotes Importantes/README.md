# Módulo 02 — Pacotes Importantes

Continuação do módulo Fundação. As pastas seguem a ordem dos 22 itens apresentados no curso.

## Estado do módulo

As aulas 01 a 21 têm exemplos implementados, com `main.go`, README, resultado esperado e desafio para aprofundamento. As aulas de arquivos estáticos e templates incluem os arquivos auxiliares necessários. O item 22 é uma referência de documentação e contém somente README.

## Aulas

1. [Manipulação de arquivos](01%20-%20Manipula%C3%A7%C3%A3o%20de%20arquivos/README.md)
2. [Realizando chamada HTTP](02%20-%20Realizando%20chamada%20HTTP/README.md)
3. [Defer](03%20-%20Defer/README.md)
4. [Trabalhando com JSON](04%20-%20Trabalhando%20com%20JSON/README.md)
5. [Busca CEP](05%20-%20Busca%20CEP/README.md)
6. [Iniciando com HTTP](06%20-%20Iniciando%20com%20HTTP/README.md)
7. [Manipulando Headers](07%20-%20Manipulando%20Headers/README.md)
8. [Criando função para BuscaCEP](08%20-%20Criando%20fun%C3%A7%C3%A3o%20para%20BuscaCEP/README.md)
9. [Finalizando resposta para o servidor HTTP](09%20-%20Finalizando%20resposta%20para%20o%20servidor%20HTTP/README.md)
10. [ServeMux](10%20-%20ServeMux/README.md)
11. [FileServer](11%20-%20FileServer/README.md)
12. [Iniciando com templates](12%20-%20Iniciando%20com%20templates/README.md)
13. [Template.Must](13%20-%20Template.Must/README.md)
14. [Templates com arquivo externo](14%20-%20Templates%20com%20arquivo%20externo/README.md)
15. [Templates com webserver](15%20-%20Templates%20com%20webserver/README.md)
16. [Compondo templates](16%20-%20Compondo%20templates/README.md)
17. [Mapeando funções nos templates](17%20-%20Mapeando%20fun%C3%A7%C3%B5es%20nos%20templates/README.md)
18. [HttpClient com Timeout](18%20-%20HttpClient%20com%20Timeout/README.md)
19. [Trabalhando com Post](19%20-%20Trabalhando%20com%20Post/README.md)
20. [Customizando Request](20%20-%20Customizando%20Request/README.md)
21. [Trabalhando com HTTP usando Contextos](21%20-%20Trabalhando%20com%20HTTP%20usando%20Contextos/README.md)
22. [Documentação para Template](22%20-%20Documenta%C3%A7%C3%A3o%20para%20Template/README.md)

## Como começar

Partindo da raiz do repositório:

```powershell
cd "02 - Pacotes Importantes\01 - Manipulação de arquivos"
go run main.go
```

Cada aula é independente e usa apenas a biblioteca padrão, sem `go.mod`. Execute `go run main.go` dentro da pasta da aula, especialmente quando houver arquivos em `public` ou `templates`.

As aulas **06, 07, 09, 10, 11 e 15** iniciam servidores em `localhost:8080`: execute uma por vez, consulte as rotas indicadas no README e encerre com Ctrl+C. As demais aulas terminam automaticamente.

As aulas 02, 05, 08, 18, 19, 20 e 21 usam servidores temporários em portas locais livres. Nas aulas de CEP (05, 08 e 09), o modo padrão simula a API sem internet; `-online` ativa a consulta real ao ViaCEP. A aula 01 usa um arquivo temporário e o remove ao terminar.

## Validar os exemplos

As aulas 06, 07, 08, 09, 10, 11, 15 e 19 incluem testes de comportamento. Dentro de uma dessas aulas, execute:

```powershell
go test main.go main_test.go
```

Esses testes usam handlers ou servidores locais para conferir respostas e erros sem depender do ViaCEP real. Não é necessário iniciar `go run` antes dos testes. Nas demais aulas, compare a execução com o resultado documentado.

## Material de apoio

- [Documentos e apostila](00%20-%20Documentos/README.md).
- [ViaCEP](https://viacep.com.br/): referência sugerida para as aulas de busca de CEP; a consulta real exige internet.
- [Skill local pacotes-importantes-go](../.github/skills/pacotes-importantes-go/SKILL.md): orientação para estudar, implementar e revisar este módulo.

[Voltar ao curso](../README.md)
