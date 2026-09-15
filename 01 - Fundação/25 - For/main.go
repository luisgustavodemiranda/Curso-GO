package main

import "fmt"

func main() {
	total := 0
	for numero := 1; numero <= 5; numero++ {
		total += numero
	}

	fmt.Println("Soma:", total)
}
