package internal

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New() (*Storage, error) {
	dbConfig, err := configNew()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", dbConfig)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS URL_S (
			original_url TEXT NOT NULL UNIQUE,
			url_to_redirect VARCHAR(50) NOT NULL UNIQUE
			);`)
	if err != nil {
		return nil, err
	}
	slog.Info("DB is create")

	return &Storage{db: db}, nil
}

func (s *Storage) SaveURL(origURL, urlToRedirect string) error {

	_, err := s.db.Exec("INSERT INTO URL_S (original_url, url_to_redirect) VALUES ($1, $2)",
		origURL, urlToRedirect)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GiveURLToRedirect(urlToRedirect string) (string, error) {

	var url string
	err := s.db.QueryRow(`SELECT original_url FROM URL_S
	 WHERE url_to_redirect = $1`, urlToRedirect).Scan(&url)
	if err != nil {
		return url, err
	}

	return url, nil
}
