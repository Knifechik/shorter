package repo

import (
	"ozontest/cmd/shorter/internal/app"
	"time"
)

type dbUrl struct {
	ID        string    `db:"id"`
	Url       string    `db:"url"`
	CreatedAt time.Time `db:"created_at"`
}

func (db dbUrl) convert() *app.MyURL {
	return &app.MyURL{
		ID:  db.ID,
		Url: db.Url,
	}
}
