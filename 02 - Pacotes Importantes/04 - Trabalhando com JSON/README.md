# Aula 04: Trabalhando com JSON

## Objetivo

Converter structs em JSON e JSON em structs com encoding/json.

## Como funciona

`json.Marshal` transforma a struct em bytes JSON. `json.Unmarshal` preenche outra struct por ponteiro. Os campos são exportados e as tags definem os nomes no JSON.

## Executar

Dentro da pasta desta aula:

```powershell
go run main.go
```

## Resultado esperado

```text
JSON: {"nome":"Ana","idade":25}
Pessoa: Ana, 25 anos
```

## Desafio prático

Adicione um campo Ativo com tag json e confirme seu valor após a decodificação.

[Voltar ao índice do módulo](../README.md)
