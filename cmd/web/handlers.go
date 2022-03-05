package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/omar-bizreh/snippetbox/pkg/models"
)

// Define a home handler function which writes a byte slice
// containing "Hello from Snippetbox" as the response body
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// s, err = app.snippets.Latest()
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

// Add a showSnippet handler function
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Fprintf(w, "%v", s)

}

// Add a createSnippet handler function
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O Snail"
	content := "O Snail content"
	expires := "7"

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
