package main

import (
	"fmt"
	"math"
)

func conditional_if() {
	if x := math.Sqrt(4); x < 1 { //short statment
		fmt.Println(x)
	} else if x > 5 {
		fmt.Println(" maior que zero")
	} else {
		fmt.Println(" caiu no else")
	}
}
