package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id, err := uuid.Parse("550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("UUID:", id)
	fmt.Println("Versao:", id.Version())
}
