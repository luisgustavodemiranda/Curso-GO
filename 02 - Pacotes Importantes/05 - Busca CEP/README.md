# Aula 05: Busca CEP

## Objetivo

Combinar cliente HTTP e JSON para consultar um CEP.

## Como funciona

A consulta é feita diretamente no `main`, combinando GET, status HTTP e `json.NewDecoder`. Por padrão, um servidor local simula o CEP `01001000`; qualquer outro CEP de oito dígitos retorna não encontrado. Use `-online` para consultar o ViaCEP real.

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

Compare os argumentos -cep 123, -cep 99999999 e o CEP padrão; observe os erros.

[Voltar ao índice do módulo](../README.md)
