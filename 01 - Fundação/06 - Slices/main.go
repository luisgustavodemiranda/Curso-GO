package main

import "fmt"

func main() {
	frutas := []string{"maca", "banana"}
	frutas = append(frutas, "laranja")

	fmt.Println("Quantidade:", len(frutas))
	fmt.Println("Frutas:", frutas)
}
