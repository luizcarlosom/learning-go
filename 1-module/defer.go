package main

import "fmt"
// Escopo do defer atrelado a função
// ela roda após o retorno da função
func doDefer() {
	defer fmt.Println(3) // LIFO
	defer fmt.Println(2)
	fmt.Println(1)
}
