package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// Server routes
func (app *application) routes() http.Handler {

	standardMiddleware := alice.New(app.recoverPanic, app.logRequests, secureHeaders)

	dynamicMiddleware := alice.New(app.session.Enable, noSurf, app.authenticate)

	mux := mux.NewRouter()
	mux.Handle("/", dynamicMiddleware.ThenFunc(app.home)).Methods(http.MethodGet)
	mux.Handle("/snippet/{id:[1-9]+}", dynamicMiddleware.ThenFunc(app.showSnippet)).Methods(http.MethodGet)

	mux.Handle("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippetForm)).Methods(http.MethodGet)
	mux.Handle("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippet)).Methods(http.MethodPost)

	mux.Handle("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm)).Methods(http.MethodGet)
	mux.Handle("/user/login", dynamicMiddleware.ThenFunc(app.loginUser)).Methods(http.MethodPost)

	mux.Handle("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm)).Methods(http.MethodGet)
	mux.Handle("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser)).Methods(http.MethodPost)

	mux.Handle("/user/logout", dynamicMiddleware.ThenFunc(app.logoutUser)).Methods(http.MethodPost)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
