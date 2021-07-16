package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *application) routes() http.Handler {
	m := mux.NewRouter()

	m.Handle("/", app.session.Enable(http.HandlerFunc(app.home))).Methods(http.MethodGet)
	m.Handle("/snippet/create", app.session.Enable(http.HandlerFunc(app.createSnippetForm))).Methods(http.MethodGet)
	m.Handle("/snippet/create", app.session.Enable(http.HandlerFunc(app.createSnippet))).Methods(http.MethodPost)
	m.Handle("/snippet/{id}", app.session.Enable(http.HandlerFunc(app.showSnippet))).Methods(http.MethodGet)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	m.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))

	middlewaredMux := app.RecoverPanic(app.LogRequest(SecureHeaders(m)))

	return middlewaredMux
}
