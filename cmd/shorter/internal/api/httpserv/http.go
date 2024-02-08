package httpserv

import (
	"github.com/gorilla/mux"
	"net/http"
	"ozontest/cmd/shorter/internal/app"
)

type application interface {
	SaveURL(*app.MyURL) (*app.MyURL, error)
	GetURL(string) (string, error)
}

type api struct {
	app application
}

func New(a application) http.Handler {
	api := api{
		app: a,
	}

	router := mux.NewRouter()

	router.HandleFunc("/", api.SaveURL).Methods(http.MethodPost)
	router.HandleFunc("/{id:[A-Za-z0-9_]+}", api.GetURL).Methods(http.MethodGet)

	return router

}
