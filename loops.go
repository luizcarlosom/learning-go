package main

import "fmt"

func loopFor() {
	var res int
	for i := 0; i < 10; i++ {
		res++
	}
	fmt.Println(res)
}

func loopForRange() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, element := range arr {
		fmt.Println(i, element)
	}
}

func loopForRangeBlankIdentify() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, element := range arr {
		fmt.Println(element)
	}
}

func loopIterandoInteiro() {
	for i := range 10 {
		fmt.Println(i)
	}
}
