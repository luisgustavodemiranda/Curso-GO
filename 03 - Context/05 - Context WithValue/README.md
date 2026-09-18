# Aula 05: Context WithValue

## Objetivo e funcionamento

WithValue associa um ID de requisição ao contexto derivado. As funções comID e idDaRequisicao encapsulam uma chave de tipo privado e a leitura com assertion segura. O filho herda o ID; o contexto original permanece sem valor. Use valores para metadados da requisição, não para parâmetros opcionais ou configuração geral.

## Executar

Requer Go instalado. Dentro da pasta desta aula:

```powershell
go run main.go
```

O programa termina sozinho, sem rede ou arquivos gerados.

## Resultado esperado

```text
[req-123] Iniciando consulta
Sem ID: Contexto original
[req-123] ID herdado
```

## Testes

```powershell
go test -timeout 20s main.go main_test.go
```

Os testes verificam comportamento, sem serviços externos. Tempos exatos dependem do agendamento do sistema.

## Desafio prático

Crie um novo contexto com ID req-456 a partir de ctx e registre os dois.

**Confira:** O novo contexto deve mostrar req-456 e o anterior deve manter req-123.

[Voltar ao módulo](../README.md)
