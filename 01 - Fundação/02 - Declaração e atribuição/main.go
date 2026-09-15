package main

func main() {
	// Declaracao explicita: a variavel existe antes de receber um valor.
	var nome string
	nome = "Maria"

	// Declaracao com valor inicial e tipo informado.
	var idade int = 20

	// Declaracao curta: o Go deduz o tipo pelo valor.
	cidade := "Sao Paulo"

	// Atribuicao: troca o valor de uma variavel ja declarada.
	idade = 21

	println(nome)
	println(idade)
	println(cidade)
}
