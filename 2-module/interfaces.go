package main

import "fmt"

type Animal interface {
	Sound() string
}

type Dog struct{}

func (Dog) Sound() string {
	return "Au Au!"
}

func AnimalFazendoZuada(a Animal) {
	fmt.Println(a.Sound())
}

func TestInterface() {
	dog := Dog{}

	AnimalFazendoZuada(dog)
}