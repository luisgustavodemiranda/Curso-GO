package main

import "fmt"

type Retangulo struct {
	Largura float64
	Altura  float64
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

func main() {
	r := Retangulo{Largura: 4, Altura: 3}
	fmt.Println("Area:", r.Area())
}
