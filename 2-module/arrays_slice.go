package main

import "fmt"

func array1() {
	arr := [10]int{}
	fmt.Println(arr)
}

func slice1() { //limite inferior sempre 0 limite superior o length do array
	slice := []int{1, 2, 3}
	fmt.Println(slice)
}

var filmes_no_db = []string{
	"bla1",
	"bla2",
	"bla3",
	"bla4",
	"bla5",
	"bla6",
	"bla7",
	"bla8",
	"bla9",
	"bla10",
}

func slice2() {
	resultsFromApi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//filmes := make([]string, 0, 10) //Se sabe o tamanho exato

	var filmes []string //Aumenta dinamicamente

	for index := range resultsFromApi {
		filme := filmes_no_db[index]
		filmes = append(filmes, filme)
		fmt.Println(cap(filmes))
	}
	fmt.Println(filmes)
}
