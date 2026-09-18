# Aula 01: Introdução aos contextos

## Objetivo e funcionamento

O exemplo cria um contexto raiz com Background e um filho com WithCancel. Mostra Err antes e depois de cancelar e aguarda Done. Cancelar o filho não cancela a raiz.

## Executar

Requer Go instalado. Dentro da pasta desta aula:

```powershell
go run main.go
```

O programa termina sozinho, sem rede ou arquivos gerados.

## Resultado esperado

```text
Raiz tem prazo: false
Antes: <nil>
Depois: context canceled
Raiz: <nil>
```

## Desafio prático

Crie dois filhos da mesma raiz e cancele somente um.

**Confira:** O outro filho e a raiz devem continuar com Err igual a nil.

[Voltar ao módulo](../README.md)
