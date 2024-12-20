package main

import "fmt"

func print_ponteiro() {
	x := 10
	p := &x

	fmt.Println(p, *p)
}

// Recebe endereço de mémoria &x e troca valor que tá no endereço pelo 100
func take(x *int ) { //TIPO PONTEIRO
	*x = 100 //VALOR DO PONTEIRO
}
