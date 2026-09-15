package main

import "fmt"

type Forma interface {
	Area() float64
}

type Quadrado struct {
	Lado float64
}

func (q Quadrado) Area() float64 {
	return q.Lado * q.Lado
}

func imprimirArea(f Forma) {
	fmt.Println("Area:", f.Area())
}

func main() {
	imprimirArea(Quadrado{Lado: 5})
}
