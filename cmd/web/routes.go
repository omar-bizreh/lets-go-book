package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// Server routes
func (app *application) routes() http.Handler {

	standardMiddleware := alice.New(app.recoverPanic, app.logRequests, secureHeaders)

	mux := mux.NewRouter()
	mux.HandleFunc("/", app.home).Methods("GET")
	mux.HandleFunc("/snippet/{id}", app.showSnippet).Methods("GET")
	mux.HandleFunc("/snippet/create", app.createSnippetForm).Methods("GET")
	mux.HandleFunc("/snippet/create", app.createSnippet).Methods("POST")

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
