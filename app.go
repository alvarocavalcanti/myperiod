package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App ...
type App struct {
	Router *mux.Router
	Entries []Entry
}

// Initialize ...
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/", HomeHandler)
	a.Router.HandleFunc("/entries", a.EntriesHandler)

	a.Entries = []Entry{}
	// entry := Entry{ ID: 1, Type: HighFlow, Date: time.Now()}
	// a.Entries = append(a.Entries, entry)
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

// EntriesHandler ...
func (a *App) EntriesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonEntries, _ := json.Marshal(a.Entries)
	io.WriteString(w, string(jsonEntries))
}
