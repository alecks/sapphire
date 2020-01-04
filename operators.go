package main

type operatorFunction func(int, int) int
type operator struct {
	Identifier string
	Function   operatorFunction
}

var operators = []operator{{"times", times}}

func times(x int, y int) int {
	return x * y
}
