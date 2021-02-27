package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	// data := []byte("Hello World!")
	data := struct{
		Message string `json:"message"`
	}{"hello world!"}
	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	// w.Write(data)
	json.NewEncoder(w).Encode(data)
}