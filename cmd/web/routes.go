package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *application) routes() http.Handler {
	m := mux.NewRouter()

	m.HandleFunc("/", app.home).Methods(http.MethodGet)
	m.HandleFunc("/snippet/create", app.createSnippetForm).Methods(http.MethodGet)
	m.HandleFunc("/snippet/create", app.createSnippet).Methods(http.MethodPost)
	m.HandleFunc("/snippet/{id}", app.showSnippet).Methods(http.MethodGet)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	m.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))

	middlewaredMux := app.RecoverPanic(app.LogRequest(SecureHeaders(m)))

	return middlewaredMux
}
