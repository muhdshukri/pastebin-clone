package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handle("http.MethodGet", "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandleFunc(http.MethodGet, "/", app.home)
	router.HandleFunc(http.MethodGet, "/snippet/view", app.snippetView)
	router.HandleFunc(http.MethodGet, "/snippet/create", app.snippetCreate)
	router.HandleFunc(http.MethodPost, "/snippet/create", app.snippetCreatePost)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	return standard.Then(router)
}
