package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {

	router := mux.NewRouter()
	// Middleware
	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(logging)
	// r.Use(limit)
	router.Use(allowCors)


	router.HandleFunc("/", app.home)

	return router
}

func allowCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
			return
		}
		next.ServeHTTP(w, r)
	})
}

type statusRecorder struct {
	http.ResponseWriter
	Status int
}


func logging(next http.Handler) http.Handler {
	logger := log.New(os.Stdout, "router |> ", log.Lshortfile|log.Ltime)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := &statusRecorder{
			ResponseWriter: w,
			Status:         200,
		}
		next.ServeHTTP(recorder, r)
		logger.Printf("Path : %s, Method: %s, status: %d\n", r.URL.Path, r.Method, recorder.Status)
	})
}
