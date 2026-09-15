package main

import "fmt"

func descrever(valor interface{}) {
	nome, ok := valor.(string)
	if ok {
		fmt.Println("Texto:", nome)
		return
	}

	fmt.Printf("Outro tipo: %v\n", valor)
}

func main() {
	descrever("Go")
	descrever(10)
}
