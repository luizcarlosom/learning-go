package main

import (
	"context"
	"net/http"
	"time"
)

func PersonalizeHttpRequests() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://www.google.com",
		nil,
	)

	if err != nil {
		panic(err)
	}

	req.Header.Set("ACCEPT", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
}
