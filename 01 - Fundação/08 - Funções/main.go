package main

import "fmt"

func somar(a, b int) int {
	return a + b
}

func main() {
	resultado := somar(4, 6)
	fmt.Println("Resultado:", resultado)
}
