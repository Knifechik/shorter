package httpserv

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"ozontest/cmd/shorter/internal/app"
)

func (a *api) SaveURL(w http.ResponseWriter, r *http.Request) {
	var url = &app.MyURL{}
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		fmt.Print(err)
		return
	}
	furl, err := a.app.SaveURL(url)
	if err != nil {
		fmt.Print(err)
	}
	err = json.NewEncoder(w).Encode(furl)
	if err != nil {
		fmt.Print(err)
	}
}

func (a *api) GetURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	url, err := a.app.GetURL(id)
	if err != nil {
		fmt.Print(err)
		return
	}
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
