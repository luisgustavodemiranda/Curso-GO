# Aula 18: HttpClient com Timeout

## Objetivo

Limitar o tempo de uma chamada com http.Client.Timeout.

## Como funciona

Um servidor local responde imediatamente em `/rapido` e aguarda 300 ms em `/lento`. O cliente tem timeout de 100 ms. O exemplo verifica um erro que implementa `net.Error` e indica timeout; o servidor temporário é encerrado automaticamente.

## Executar

Dentro da pasta desta aula:

```powershell
go run main.go
```

## Resultado esperado

```text
Resposta rapida: OK
Resposta lenta: timeout
```

## Desafio prático

Aumente o timeout acima de 300 ms e adapte a verificação para esperar sucesso na resposta lenta.

[Voltar ao índice do módulo](../README.md)
