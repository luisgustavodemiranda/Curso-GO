package main

import "fmt"

func criarContador() func() int {
	contador := 0

	return func() int {
		contador++
		return contador
	}
}

func main() {
	contador := criarContador()
	fmt.Println(contador())
	fmt.Println(contador())
}
