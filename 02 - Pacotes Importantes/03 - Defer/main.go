package main

import "fmt"

func exemplo() {
	fmt.Println("inicio")
	// Os defers executam quando esta função retorna, do último para o primeiro.
	defer fmt.Println("primeiro defer")
	defer fmt.Println("segundo defer")
	fmt.Println("fim")
}

func main() { exemplo() }
