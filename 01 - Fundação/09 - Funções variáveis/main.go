package main

import "fmt"

func somar(valores ...int) int {
	total := 0
	for _, valor := range valores {
		total += valor
	}
	return total
}

func main() {
	fmt.Println(somar(1, 2, 3, 4))
}
