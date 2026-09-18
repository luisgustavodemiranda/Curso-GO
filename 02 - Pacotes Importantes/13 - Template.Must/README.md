# Aula 13: Template.Must

## Objetivo

Usar template.Must na inicialização de templates conhecidos.

## Como funciona

`template.Must` recebe o resultado de Parse e provoca panic se houver erro. É usado aqui com um template fixo na inicialização. O erro de Execute continua sendo tratado normalmente.

## Executar

Dentro da pasta desta aula:

```powershell
go run main.go
```

## Resultado esperado

```text
Ola, Ana!
```

## Desafio prático

Introduza uma ação de template incompleta para observar o panic; depois restaure.

[Voltar ao índice do módulo](../README.md)
