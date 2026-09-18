# Aula 03: Context utilizando server HTTP

## Objetivo e funcionamento

O handler deriva um prazo de 100 ms de r.Context e simula uma consulta de um segundo. O caminho / responde 504 quando o prazo local termina; /?rapido=1 conclui imediatamente com 200. Se o contexto original da requisição for cancelado, o handler abandona o trabalho. O status 504 é uma escolha explícita deste exemplo, não algo enviado automaticamente por context.

## Executar

Requer Go instalado. Dentro da pasta desta aula:

```powershell
go run main.go
```

O servidor permanece em `localhost:8080`. Em outro terminal:

```powershell
curl.exe -i "http://localhost:8080/"
curl.exe -i "http://localhost:8080/?rapido=1"
```

Encerre com **Ctrl+C**. Execute somente uma aula que use essa porta por vez. Para observar uma desconexão, use `curl.exe --max-time 0.05 "http://localhost:8080/"`: o cliente deve encerrar antes do prazo do servidor, que registra o cancelamento.

## Resultado esperado

```text
/ → HTTP 504: Prazo do servidor esgotado
/?rapido=1 → HTTP 200: Consulta concluida
```

## Testes

```powershell
go test -timeout 20s main.go main_test.go
```

Os testes verificam comportamento, sem serviços externos. Tempos exatos dependem do agendamento do sistema.

## Desafio prático

Aumente o prazo do handler para dois segundos e consulte / novamente.

**Confira:** A resposta deve mudar de 504 para 200 com Consulta concluida.

[Voltar ao módulo](../README.md)
