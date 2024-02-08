package repo

import (
	"context"
	"database/sql"
	"fmt"
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

func (r *Repo) Insert(ctx context.Context, u *app.MyURL) error {
	const query = `INSERT INTO url_table (id, url) VALUES ($1, $2)`
	_, err := r.sql.ExecContext(ctx, query, u.ID, u.Url)
	if err != nil {
		return fmt.Errorf("\nsql.ExecContext: %w", err)
	}
	return nil
}

func (r *Repo) GetURL(ctx context.Context, id string) (*app.MyURL, error) {
	url := dbUrl{}
	const query = `SELECT url FROM url_table WHERE id = $1`
	err := r.sql.QueryRowContext(ctx, query, id).Scan(&url.Url)
	if err != nil {
		return nil, fmt.Errorf("sql.QueryRowContext: %w", err)
	}
	res := url.convert()

	return res, nil
}
