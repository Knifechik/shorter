package app

import "context"

type Repo interface {
	Insert(context.Context, *MyURL) error
	GetURL(context.Context, string) (*MyURL, error)
}
