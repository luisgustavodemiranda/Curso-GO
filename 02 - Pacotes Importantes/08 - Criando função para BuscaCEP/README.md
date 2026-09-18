# Aula 08: Criando função para BuscaCEP

## Objetivo

Separar a consulta de CEP em uma função reutilizável.

## Como funciona

`buscarCEP` recebe cliente, URL base e CEP, devolvendo `Endereco` e `error`. Isso separa a consulta da apresentação e permite usar a mesma função com o servidor local ou ViaCEP. Há erros distintos para formato inválido, CEP inexistente, status inesperado e JSON inválido.

## Executar

Dentro da pasta desta aula:

```powershell
go run main.go
```

## Consulta real de CEP

O modo padrão funciona sem internet e reconhece apenas o CEP `01001000`. Para consultar o [ViaCEP](https://viacep.com.br/) real:

```powershell
go run main.go -online -cep 01001000
```

A chamada externa tem timeout de cinco segundos. Os dados reais podem conter acentos e diferir da amostra local. As aulas 05 e 08 encerram com erro para CEP inválido/inexistente; a aula 09 comunica o erro pelo status HTTP.

## Resultado esperado

```text
01001-000: Praca da Se, Sao Paulo/SP
```

## Desafio prático

Chame buscarCEP duas vezes, com um CEP conhecido e um inexistente, tratando os erros no chamador.

[Voltar ao índice do módulo](../README.md)
