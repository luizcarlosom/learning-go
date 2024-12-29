package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	respTwo, err := http.Post("https://minhaapi.com", "application/json", nil)

	if err != nil {
		panic(err)
	}

	defer respTwo.Body.Close()

	dataTwo, err := io.ReadAll(respTwo.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataTwo))
}
