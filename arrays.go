package main

import "fmt"

func array1() {
	arr := [3]int{}
	fmt.Println(arr)
}

func array10() {
	arr := [10]int{5: 400, 7: 300}
	fmt.Println(arr)
}

func array11() {
	arr := [10]int{5}
	fmt.Println(arr)
}