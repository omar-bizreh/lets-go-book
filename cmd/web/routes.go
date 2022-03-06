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
	mux.HandleFunc("/", app.home).Methods(http.MethodGet)
	mux.HandleFunc("/snippet/{id:[1-9]+}", app.showSnippet).Methods(http.MethodGet)

	mux.HandleFunc("/snippet/create", app.createSnippetForm).Methods(http.MethodGet)
	mux.HandleFunc("/snippet/create", app.createSnippet).Methods(http.MethodPost)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
