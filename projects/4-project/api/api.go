package api

import (
	"4-project/omdb"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func NewHandler(apiKey string) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/", handleSearchMovie(apiKey))

	return r
}

func sendJson(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("Failed to marshal json data", "error", err)
		sendJson(
			w,
			Response{Error: "something went wrong"},
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("Failed to qrite response to cliente", "error", err)
		return
	}
}

func handleSearchMovie(apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("s")
		res, err := omdb.Search(apiKey, search)

		if err != nil {
			sendJson(
				w,
				Response{Error: "Something wrong with omdb"},
				http.StatusBadGateway,
			)
			return
		}
		sendJson(w, Response{Data: res}, http.StatusOK)
	}
}
