package models

import "github.com/whoiswentz/snippetbox/pkg/forms"

type TemplateData struct {
	CurrentYear     int
	Form            *forms.Form
	Snippet         *Snippet
	Snippets        []*Snippet
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}
