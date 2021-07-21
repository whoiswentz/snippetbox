package main

import (
	"bytes"
	"fmt"
	models "github.com/whoiswentz/snippetbox/pkg"
	"net/http"
	"runtime/debug"
	"time"
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

func (app *application) addDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	if td == nil {
		td = &models.TemplateData{}
	}
	td.CurrentYear = time.Now().Year()
	td.Flash = app.session.PopString(r, "flash")
	return td
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *models.TemplateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("the template %s does not exists", name))
		return
	}

	buf := new(bytes.Buffer)
	if err := ts.Execute(buf, app.addDefaultData(td, r)); err != nil {
		app.serverError(w, err)
		return
	}

	if _, err := buf.WriteTo(w); err != nil {
		app.serverError(w, err)
	}
}
