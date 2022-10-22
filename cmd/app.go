package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Arcady1/Avito-2022/pkg/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) InitDB(user, password, dbname string) {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbname)

	database, err := sql.Open("postgres", dbInfo)

	if err != nil {
		err = errors.New("Error: connecting to the Database")
		log.Fatal(err)
	} else {
		a.DB = database
	}
}

func (a *App) InitRoutes() {
	a.Router = mux.NewRouter()

	a.Router.HandleFunc("/api/v1/user/balance", handlers.GetUserBalance).Methods("GET")
}

func (a *App) Run(host, port string) {
	http.ListenAndServe(host+":"+port, a.Router)
}
