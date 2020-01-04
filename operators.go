package main

type operatorFunction func(int, int) int
type operator struct {
	Identifier string
	Function   operatorFunction
}

var operators = []operator{{"multiply", multiply}, {"divide", divide}, {"add", add}, {"subtract", subtract}}

func multiply(x int, y int) int {
	return x * y
}

func divide(x int, y int) int {
	return x / y
}

func add(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}
