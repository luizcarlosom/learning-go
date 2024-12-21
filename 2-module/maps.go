package main

import "fmt"

func map1() {
	m := map[string]string{
		"Pedro": "Pessoa",
		"Joaquim": "Pedro",
	}

	fmt.Println(m["Pedro"])
	delete(m, "Pedro")
	fmt.Println(m["Pedro"])
}

func map2() {
	m := make(map[string]string)
	m["Pedro"] = "Pessoa"
	valor := m["Pedro"]
	fmt.Println(valor)
	delete(m, "Pedro")
	fmt.Println(m["Pedro"])
}