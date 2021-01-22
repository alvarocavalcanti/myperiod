package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App ...
type App struct {
	Router *mux.Router
	DB *sql.DB
}

// Initialize ...
func (a *App) Initialize(user, password, dbname string) { }

// Run ...
func (a *App) Run(addr string) { }
