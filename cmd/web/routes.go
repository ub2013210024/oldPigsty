package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// Create a multiplexer
	router := httprouter.New()
	// Create a file server
	fileServer := http.FileServer(http.Dir("."))

	router.Handler(http.MethodGet, "", http.StripPrefix("/static", fileServer))
	// -> Home
	router.HandlerFunc(http.MethodGet, "/", app.Home)
	// -> another handler

	return router
}
