package httpserv

import (
	"context"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"ozontest/cmd/shorter/internal/app"
)

type application interface {
	SaveURL(context.Context, *app.MyURL) (*app.MyURL, error)
	GetURL(context.Context, string) (*app.MyURL, error)
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

func errconv(w http.ResponseWriter, err error) {
	slog.Error("ERROR:", err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
