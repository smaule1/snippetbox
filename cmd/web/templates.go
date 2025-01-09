package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"snippetbox.samuel/internal/models"
	"snippetbox.samuel/ui"
)

type templateData struct {
	Snippet         models.Snippet
	Snippets        []models.Snippet
	CurrentYear     int
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
	User            models.User
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.tmpl.html",
			"html/partials/*.html",
			page,
		}

		ts, err := template.New(name).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

// Custom functions in the template
// I will not add them to the code because they seem to be useless, its way easier
// to use the function before adding the data to the template instead of doing it in the
// template parsing process
// faaaaaalse
// they are actually useful
// if we dont use them we would have to add a struct or variable for every modified data
// that we want to add
// BUuut i will not implement it because all this server-side frontend is too cumbersome in go,
// i will use a frontend-js framework if i need to
func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.UTC().Format("02 Jan 2006 at 15:04")
}

// var functions = template.FuncMap{
//     "humanDate": humanDate,
// }
// ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl")
