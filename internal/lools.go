package internal

import (
	"crypto/md5"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"try/internal/storage"

	"github.com/gorilla/mux"
)

type URLShortener struct {
	storage *storage.Storage
}

func NewURLShortener(s *storage.Storage) *URLShortener {
	return &URLShortener{storage: s}
}

func (us *URLShortener) Shorten(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.Form.Get("url")
	if !(strings.Contains(url, "http://") || strings.Contains(url, "https://")) {
		fmt.Println("Please give normal link")
	} else {
		shortURL := fmt.Sprintf("%x", md5.Sum([]byte(url)))[:5]
		us.storage.SaveURL(url, shortURL)
		fmt.Printf("%s/%s\n", os.Getenv("HOST"), shortURL)
	}
}

func (us *URLShortener) CreateRedirectLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	originalURL, err := us.storage.GiveURLToRedirect(vars["shortURL"])
	if err != nil {
		slog.Any("ERROR", err)
		return
	}
	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}
