package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {

	router := mux.NewRouter()

	router.HandleFunc("/", app.home)

	return router
}