# Item 20: Generics

Generics permitem escrever uma função que trabalha com mais de um tipo mantendo a segurança de tipos. Em vez de criar uma função diferente para `int` e outra para `float64`, podemos criar uma única função e indicar quais tipos ela aceita.

## Restrições de tipos

```go
type Number interface {
	~int | ~float64
}
```

`Number` é uma restrição: ela informa que o tipo genérico pode ser `int` ou `float64`. O símbolo `~` também permite tipos definidos pelo usuário cujo tipo subjacente seja um desses tipos:

```go
type MyNumber int
```

Por isso, `MyNumber` pode ser usado com `maior`:

```go
maior(MyNumber(10), MyNumber(20)) // resultado: 20
```

## Funções genéricas

A função `maior` recebe dois valores do mesmo tipo `T` e devolve o maior deles:

```go
func maior[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}
```

Nas chamadas abaixo, o Go infere o tipo `T` pelos argumentos. Assim, a mesma função trabalha com inteiros, números decimais e `MyNumber`:

```go
maior(3, 7)                         // 7
maior(2.5, 1.8)                     // 2.5
maior(MyNumber(10), MyNumber(20))   // 20
```

A restrição `comparable` permite usar os operadores `==` e `!=` com o tipo genérico. Por isso, `igual` pode comparar valores de tipos comparáveis, como `string`, `int` e `float64`:

```go
func igual[T comparable](a, b T) bool {
	return a == b
}

igual("Go", "Go") // true
```

Ela é diferente de `Number`: `comparable` serve para igualdade, enquanto `Number` limita os tipos que podem ser usados com o operador `>`.

A função `Soma` também usa a restrição `Number`. Ela percorre os dois argumentos, soma seus valores e retorna o resultado no mesmo tipo `T`:

```go
Soma(3, 7) // 10
```

## Saída do programa

Ao executar o arquivo, as chamadas do `main` produzem:

```text
7
2.5
20
true
10
```

## Executar

```powershell
go run main.go
```
