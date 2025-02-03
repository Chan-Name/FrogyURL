package internal

import (
	"crypto/md5"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func (s *Storage) Shorten(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.Form.Get("url")
	if !(strings.Contains(url, "http://") || strings.Contains(url, "https://")) {
		fmt.Println("Please give normal link")
	} else {
		shortURL := fmt.Sprintf("%x", md5.Sum([]byte(url)))[:5]
		s.SaveURL(url, shortURL)
		fmt.Printf("http://localhost:8080/%s\n", shortURL)
	}
}

func (s *Storage) CreateRedirectLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	originalURL, err := s.GiveURLToRedirect(vars["shortURL"])
	if err != nil {
		slog.Any("ERROR", err)
		return
	}
	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}
