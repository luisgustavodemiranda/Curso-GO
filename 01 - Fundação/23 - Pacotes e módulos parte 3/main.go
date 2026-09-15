package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	caminho := filepath.Join("dados", "usuarios.txt")
	fmt.Println(caminho)
}
