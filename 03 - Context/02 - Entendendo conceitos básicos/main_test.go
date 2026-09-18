package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestTrabalhar(t *testing.T) {
	if err := trabalhar(context.Background(), 0); err != nil {
		t.Fatal(err)
	}
	ctx, cancelar := context.WithCancel(context.Background())
	cancelar()
	if err := trabalhar(ctx, time.Hour); !errors.Is(err, context.Canceled) {
		t.Fatalf("cancelamento: %v", err)
	}
	prazo, liberar := context.WithDeadline(context.Background(), time.Now().Add(-time.Second))
	defer liberar()
	if err := trabalhar(prazo, time.Hour); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("prazo: %v", err)
	}
}
