package httpserv

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"ozontest/cmd/shorter/internal/app"
)

func (a *api) SaveURL(w http.ResponseWriter, r *http.Request) {
	var url = &app.MyURL{}
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		errconv(w, err)
		return
	}

	furl, err := a.app.SaveURL(r.Context(), url)
	if err != nil {
		errconv(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(furl)
	if err != nil {
		errconv(w, err)
		return
	}
}

func (a *api) GetURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	url, err := a.app.GetURL(r.Context(), id)
	if err != nil {
		errconv(w, err)
		return
	}

	http.Redirect(w, r, url.Url, http.StatusMovedPermanently)
}
