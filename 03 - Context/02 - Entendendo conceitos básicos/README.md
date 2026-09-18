# Aula 02: Entendendo conceitos básicos

## Objetivo e funcionamento

A função trabalhar recebe ctx como primeiro argumento e usa select para observar Done durante a espera. O main compara sucesso, cancelamento explícito, timeout e cancelamento herdado do pai. Cancelamento é cooperativo: não interrompe automaticamente código que ignora o contexto. WithDeadline usa um instante absoluto; WithTimeout recebe uma duração.

## Executar

Requer Go instalado. Dentro da pasta desta aula:

```powershell
go run main.go
```

O programa termina sozinho, sem rede ou arquivos gerados.

## Resultado esperado

```text
Concluido: <nil>
Cancelado: context canceled
Timeout: context deadline exceeded
Filho: context canceled
```

## Testes

```powershell
go test -timeout 20s main.go main_test.go
```

Os testes verificam comportamento, sem serviços externos. Tempos exatos dependem do agendamento do sistema.

## Desafio prático

Substitua WithTimeout por WithDeadline com time.Now().Add(50*time.Millisecond).

**Confira:** O trabalho lento deve continuar retornando context deadline exceeded.

[Voltar ao módulo](../README.md)
