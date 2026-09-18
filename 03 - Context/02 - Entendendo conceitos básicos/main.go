package main

import (
	"context"
	"fmt"
	"time"
)

func trabalhar(ctx context.Context, duracao time.Duration) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	temporizador := time.NewTimer(duracao)
	defer temporizador.Stop()
	// O cancelamento é cooperativo: o trabalho precisa observar Done.
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-temporizador.C:
		return nil
	}
}
func main() {
	fmt.Println("Concluido:", trabalhar(context.Background(), 0))
	cancelado, cancelar := context.WithCancel(context.Background())
	cancelar()
	fmt.Println("Cancelado:", trabalhar(cancelado, time.Second))
	ctx, liberar := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer liberar()
	fmt.Println("Timeout:", trabalhar(ctx, time.Second))
	pai, cancelarPai := context.WithCancel(context.Background())
	defer cancelarPai()
	filho, cancelarFilho := context.WithCancel(pai)
	defer cancelarFilho()
	cancelarPai()
	fmt.Println("Filho:", filho.Err())
}
