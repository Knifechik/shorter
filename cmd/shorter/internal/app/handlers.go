package app

import (
	"context"
	"fmt"
)

func (a *App) SaveURL(ctx context.Context, url *MyURL) (*MyURL, error) {
	url.ID = RandomShortURL()
	err := a.repo.Insert(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("repo.Insert: %w", err)
	}
	return url, nil
}

func (a *App) GetURL(ctx context.Context, id string) (*MyURL, error) {
	url, err := a.repo.GetURL(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("repo.GetUrl: %w", err)
	}
	return url, nil
}
