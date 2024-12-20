package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)
}

func create() *int {
	x := 10
	return &x
}