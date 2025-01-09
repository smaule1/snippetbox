package main

import (
	"net/http"

	"github.com/justinas/alice"
	"snippetbox.samuel/ui"
)

func (app *application) routes() http.Handler {
	// Servemux connects the routes with the handlers
	mux := http.NewServeMux()

	// File server for static files
	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	//Simple ping
	mux.HandleFunc("GET /ping", ping)

	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)
	// Web server routes
	// unprotected
	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetView))
	mux.Handle("GET /user/signup", dynamic.ThenFunc(app.userSignup))
	mux.Handle("POST /user/signup", dynamic.ThenFunc(app.userSignupPost))
	mux.Handle("GET /user/login", dynamic.ThenFunc(app.userLogin))
	mux.Handle("POST /user/login", dynamic.ThenFunc(app.userLoginPost))
	mux.Handle("GET /about", dynamic.ThenFunc(app.about))

	protected := dynamic.Append(app.requireAuthentication)
	//protected
	mux.Handle("GET /snippet/create", protected.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", protected.ThenFunc(app.snippetCreatePost))
	mux.Handle("POST /user/logout", protected.ThenFunc(app.userLogoutPost))
	mux.Handle("GET /account/view", protected.ThenFunc(app.account))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)
	return standard.Then(mux)
}
