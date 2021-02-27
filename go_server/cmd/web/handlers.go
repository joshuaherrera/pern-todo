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
	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (app *application) getTodos(w http.ResponseWriter, r *http.Request) {
	t, err := app.todos.All()
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.infoLog.Println("**** queried postgres")
	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}