package app

func (a *App) SaveURL(url *MyURL) (*MyURL, error) {
	url.ID = RandomShortURL()
	err := a.repo.Insert(url)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (a *App) GetURL(id string) (string, error) {
	url, err := a.repo.GetURL(id)
	if err != nil {
		return "", err
	}
	return url, nil
}
