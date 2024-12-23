package main

import (
	"errors"
	"fmt"
	"math"
)

func error1(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Não pode dividir nenhum número por 0.")
	}
	return a / b, nil
}

type SqrtError struct {
	msg string
}

func (s SqrtError) Error() string { return s.msg }

func RaizQuadrada(x float64) (float64, error) {
	if x < 0 {
		return 0, SqrtError{"Não existe raiz quadrada para número negativo"}
	}
	resultado := math.Sqrt(x)
	return resultado, nil
}

func MyError() {
	x := 10

	res, err := RaizQuadrada(float64(x))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
