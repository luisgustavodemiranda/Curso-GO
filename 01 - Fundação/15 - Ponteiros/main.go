package main

import "fmt"

func dobrar(valor *int) {
	*valor *= 2
}

func main() {
	numero := 10
	dobrar(&numero)
	fmt.Println(numero)
}
