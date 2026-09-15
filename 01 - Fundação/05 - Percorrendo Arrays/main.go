package main

import "fmt"

func main() {
	notas := [4]float64{8.5, 7.0, 9.0, 10.0}

	for indice, nota := range notas {
		fmt.Printf("Nota %d: %.1f\n", indice+1, nota)
	}
}
