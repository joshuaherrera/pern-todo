package main

import "net/http"

func (app *application) home(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	data := []byte("Hello World!")

	w.Write(data)
}