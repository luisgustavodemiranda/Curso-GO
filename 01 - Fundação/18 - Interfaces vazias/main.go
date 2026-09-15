package main

import "fmt"

func exibir(valor interface{}) {
	fmt.Printf("Valor: %v\n", valor)
}

func main() {
	exibir("texto")
	exibir(42)
	exibir(true)
}
