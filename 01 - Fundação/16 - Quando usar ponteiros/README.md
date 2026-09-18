# Item 16: quando usar ponteiros

Ponteiros são úteis quando uma função precisa alterar o valor original ou quando se deseja evitar a cópia de uma struct grande.

O exemplo atualiza uma configuração recebida por ponteiro.

## Executar

```powershell
go run main.go
```

## Desafio prático

Compare a função atual com outra que receba `Configuracao` por valor e altere `Nome`. Execute cada uma sobre uma configuração nova.

**Confira:** A versão por valor mantém o original como `teste`; a versão por ponteiro muda o original para `producao`.
