package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var hashParser = map[string]string{}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/parser", createHashParser)
	r.Get("/{hash}", getLinkFromHash)
	http.ListenAndServe(":8080", r)
}

func createHashParser(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	t := struct {
		Url string `json:"url"`
	}{}
	err := d.Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashParser["123"] = t.Url
	w.Write([]byte(t.Url))
}

func getLinkFromHash(w http.ResponseWriter, r *http.Request) {
	hash := chi.URLParam(r, "hash")
	url, ok := hashParser[hash]
	if ok {
		http.Redirect(w, r, url, http.StatusSeeOther)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not found!"))
	}
}
