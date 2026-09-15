# Item 24: instalando pacotes

O exemplo usa o pacote padrão `net/http`, que já faz parte do Go. Pacotes externos normalmente são adicionados com `go get caminho/do/pacote`.

Depois de importar um pacote, `go mod tidy` pode atualizar as dependências do módulo.

## Executar

```powershell
go run main.go
```
