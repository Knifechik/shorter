package app

type Repo interface {
	Insert(*MyURL) error
	GetURL(string) (string, error)
}
