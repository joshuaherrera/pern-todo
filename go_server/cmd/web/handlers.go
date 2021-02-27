package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	// data := []byte(`{"mesage":"hello world!"}`)
	// w.Write(data)

	data := struct{
		Message string `json:"message"`
	}{"hello world!"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(data)
}

func (app *application) getTodos(w http.ResponseWriter, r *http.Request) {
	t, err := app.todos.All()
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(t)
}