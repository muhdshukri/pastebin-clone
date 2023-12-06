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

	router.HandleFunc("/", app.home)
	router.HandleFunc("/snippet/view", app.snippetView)
	router.HandleFunc("/snippet/create", app.snippetCreate)
	router.HandleFunc("/snippet/create", app.SnippetCreatePost)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	return standard.Then(router)
}
