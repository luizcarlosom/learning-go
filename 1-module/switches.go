package main

import "fmt"

func do(x int) {
	switch x {
	case 1:
		fmt.Println("Um")
	case 2, 3:
		fmt.Println("Dois ou Três")
	default:
		fmt.Println("outra coisa")
	}
}
