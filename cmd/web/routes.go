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

	m.Handle("/user/signup", app.session.Enable(http.HandlerFunc(app.signupUserForm))).Methods(http.MethodGet)
	m.Handle("/user/signup", app.session.Enable(http.HandlerFunc(app.signupUser))).Methods(http.MethodPost)
	m.Handle("/user/login", app.session.Enable(http.HandlerFunc(app.loginUserForm))).Methods(http.MethodGet)
	m.Handle("/user/login", app.session.Enable(http.HandlerFunc(app.loginUser))).Methods(http.MethodPost)
	m.Handle("/user/logout", app.session.Enable(http.HandlerFunc(app.logoutUser))).Methods(http.MethodPost)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	m.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))

	middlewaredMux := app.RecoverPanic(app.LogRequest(SecureHeaders(m)))

	return middlewaredMux
}
