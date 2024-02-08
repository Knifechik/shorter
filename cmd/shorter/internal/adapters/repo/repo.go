package repo

import (
	"database/sql"
	_ "github.com/lib/pq"
	"ozontest/cmd/shorter/internal/app"
)

var _ app.Repo = &Repo{}

type Repo struct {
	sql *sql.DB
}

func New(name string, info string) (*Repo, error) {
	db, err := sql.Open(name, info)
	if err != nil {
		return nil, err
	}
	return &Repo{db}, err
}

func (r *Repo) Close() error {
	return r.sql.Close()
}

func (r *Repo) Insert(u *app.MyURL) error {
	const query = `INSERT INTO url_table (id, url) VALUES ($1, $2)`
	_, err := r.sql.Exec(query, u.ID, u.Url)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetURL(id string) (string, error) {
	var url string
	const query = `SELECT url FROM url_table WHERE id = $1`
	err := r.sql.QueryRow(query, id).Scan(&url)
	if err != nil {
		return "", err
	}
	return url, nil
}
