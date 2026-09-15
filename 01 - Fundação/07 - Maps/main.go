package main

import "fmt"

func main() {
	idades := map[string]int{
		"Ana":    25,
		"Carlos": 31,
	}

	idades["Maria"] = 20
	idade, existe := idades["Ana"]

	fmt.Println("Ana:", idade, "Existe:", existe)
	fmt.Println("Mapa:", idades)
}
