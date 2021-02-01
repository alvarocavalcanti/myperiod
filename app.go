package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App ...
type App struct {
	Router *mux.Router
}

// Initialize ...
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/", HomeHandler)
}

// Run ...
func (a *App) Run() {
	log.Fatal(http.ListenAndServe("localhost:8080", a.Router))
}

// HomeHandler ...
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{}`)
}
