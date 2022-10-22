package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Arcady1/go-rest-api/pkg/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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
		log.Fatalln(err)
	} else {
		a.DB = database
	}
}

func (a *App) InitRoutes() {
	a.Router = mux.NewRouter()

	a.Router.HandleFunc("/api/v1/user/balance", handlers.GetUserBalance).Methods(http.MethodPost)
	a.Router.HandleFunc("/api/v1/user/balance", handlers.IncreaseUserBalance).Methods(http.MethodPut)
}

func (a *App) Run(host, port string) {
	listenAdress := fmt.Sprintf("%s:%s", host, port)

	fmt.Println("Started on: http://" + listenAdress)

	http.ListenAndServe(listenAdress, a.Router)

}

func main() {
	a := App{}

	// Load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("Error: loading .env file")
	}

	// Inits
	a.InitDB(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.InitRoutes()

	a.Run(
		os.Getenv("HOST"),
		os.Getenv("PORT"))
}
