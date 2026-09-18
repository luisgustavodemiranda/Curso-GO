package main

import (
	"context"
	"fmt"
)

// Tipo privado evita colisões com chaves de outros pacotes.
type chaveRequisicao struct{}

func comID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, chaveRequisicao{}, id)
}
func idDaRequisicao(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(chaveRequisicao{}).(string)
	return id, ok
}
func registrar(ctx context.Context, mensagem string) {
	id, ok := idDaRequisicao(ctx)
	if !ok {
		fmt.Println("Sem ID:", mensagem)
		return
	}
	fmt.Printf("[%s] %s\n", id, mensagem)
}
func main() {
	raiz := context.Background()
	ctx := comID(raiz, "req-123")
	registrar(ctx, "Iniciando consulta")
	registrar(raiz, "Contexto original")
	filho, cancelar := context.WithCancel(ctx)
	defer cancelar()
	registrar(filho, "ID herdado")
}
