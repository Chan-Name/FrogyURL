package main

import (
	"log/slog"
	"net/http"
	"try/internal"
	"try/internal/storage"

	"github.com/gorilla/mux"
)

func main() {
	db, err := storage.New()
	if err != nil {
		slog.Any("ERROR", err)
	}

	a := internal.NewURLShortener(db)

	r := mux.NewRouter()
	r.HandleFunc("/create", a.Shorten).Methods("POST")
	r.HandleFunc("/{shortURL}", a.CreateRedirectLink).Methods("GET")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		slog.Any("ERROR", err)
		return
	}

}
