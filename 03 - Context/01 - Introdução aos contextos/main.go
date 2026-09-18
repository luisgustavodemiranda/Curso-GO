package main

import (
	"context"
	"fmt"
)

func main() {
	raiz := context.Background()
	_, temPrazo := raiz.Deadline()
	fmt.Println("Raiz tem prazo:", temPrazo)
	ctx, cancelar := context.WithCancel(raiz)
	defer cancelar()
	fmt.Println("Antes:", ctx.Err())
	cancelar()
	<-ctx.Done()
	fmt.Println("Depois:", ctx.Err())
	fmt.Println("Raiz:", raiz.Err())
}
