package main

import (
	"log/slog"
	"net/http"
	"try/internal"

	"github.com/gorilla/mux"
)

func main() {
	db, err := internal.New()
	if err != nil {
		slog.Any("ERROR", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/create", db.Shorten).Methods("POST")
	r.HandleFunc("/{shortURL}", db.CreateRedirectLink).Methods("GET")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		slog.Any("ERROR", err)
		return
	}

}
