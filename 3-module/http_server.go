package main

import (
	"errors"
	"fmt"
	"net/http"
)

func HttpServer() {
	mux := http.NewServeMux()

	mux.HandleFunc(
		"/ping",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "pong")
		},
	)

	srv := &http.Server{
		Addr:                         ":8080",
		Handler:                      mux,
		DisableGeneralOptionsHandler: false,
	}

	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}

	fmt.Println(1)
}
