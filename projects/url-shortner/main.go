package main

import (
	"log/slog"
	"net/http"
	"time"
	"url-shortner/api"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute code", "error", err)
	}
}

func run() error {
	db := make(map[string]string)


	handler := api.NewHandler(db)

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
