package main

import "fmt"

type MyNumber int

type Number interface {
	~int | ~float64
}

func maior[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func igual[T comparable](a, b T) bool {
	return a == b
}

func Soma[T Number](a, b T) T {
	var soma T
	for _, v := range []T{a, b} {
		soma += v
	}
	return soma
}

func main() {
	fmt.Println(maior(3, 7))
	fmt.Println(maior(2.5, 1.8))
	fmt.Println(maior(MyNumber(10), MyNumber(20)))
	fmt.Println(igual("Go", "Go"))
	fmt.Println(Soma(3, 7))
}
