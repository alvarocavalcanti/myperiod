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
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": ping}`)
}
