package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
	highOrder := soma1(10)(10)
	fmt.Println(highOrder)

	variadica := soma2(10, 10, 10)
	fmt.Println(variadica)

	n1, n2 := testA(50, 100)
	fmt.Println(n1, n2)

	value, module := div(10, 3)
	fmt.Println(value, module)

	value2, module2 := div2(10, 3)
	fmt.Println(value2, module2)
}

// Função High Order
func soma1(n1 int) func(int) int {
	return func(n2 int) int {
		// Função Close
		return n1 + n2
	}
}

// Função Variadica
func soma2(nums ...int) int {
	var out int
	for _, n := range nums {
		out += n
	}
	return out
}

// Função com dois retornos
func testA(n1, n2 int) (int, int) {
	return n1, n2
}

// Função com dois retornos
func div(n1, n2 int) (int, int) {
	value := n1 / n2
	module := n1 % n2
	return value, module
}

// Função pelada (não é muito utilizada)
func div2(n1, n2 int) (value, module int) {
	value = n1 / n2
	module = n1 % n2
	return
}
