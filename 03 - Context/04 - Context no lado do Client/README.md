# Aula 04: Context no lado do Client

## Objetivo e funcionamento

A função buscar recebe um contexto e cria uma requisição com NewRequestWithContext. Um servidor temporário local responde rapidamente ou espera. O exemplo mostra sucesso e um prazo de 100 ms expirando antes da resposta lenta. Client.Timeout é um limite adicional de três segundos; errors.Is identifica o erro do contexto. O corpo e o servidor são fechados antes de tratar erros no main.

## Executar

Requer Go instalado. Dentro da pasta desta aula:

```powershell
go run main.go
```

O programa abre uma porta local livre temporariamente e termina sozinho, sem internet ou dependência da aula 03.

## Resultado esperado

```text
Resposta recebida
Requisicao cancelada: prazo do contexto esgotado
```

## Testes

```powershell
go test -timeout 20s main.go main_test.go
```

Os testes verificam comportamento, sem serviços externos. Tempos exatos dependem do agendamento do sistema.

## Desafio prático

Use um contexto já cancelado em buscar, antes de enviar a requisição.

**Confira:** errors.Is(err, context.Canceled) deve ser true.

[Voltar ao módulo](../README.md)
