package main

import (
	"fmt"
	models "github.com/whoiswentz/snippetbox/pkg"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errLog.Output(2, trace)
	code := http.StatusInternalServerError
	http.Error(w, http.StatusText(code), code)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *models.TemplateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("the template %s does not exists", name))
		return
	}

	if err := ts.Execute(w, td); err != nil {
		app.serverError(w, err)
	}
}