package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Arcady1/go-rest-api/pkg/handlers"
	"github.com/Arcady1/go-rest-api/pkg/models"
	"github.com/Arcady1/go-rest-api/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type App struct {
	Router *mux.Router
}

func (a *App) InitRoutes() {
	a.Router = mux.NewRouter()

	a.Router.HandleFunc("/api/v1/user/balance", handlers.GetUserBalance).Methods(http.MethodPost)
	a.Router.HandleFunc("/api/v1/user/refill", handlers.RefillUserAccount).Methods(http.MethodPut)
	a.Router.HandleFunc("/api/v1/payments/reserve", handlers.ReserveUsersAccountMoney).Methods(http.MethodPost)
	a.Router.HandleFunc("/api/v1/payments/accept", handlers.AcceptReservedMoney).Methods(http.MethodPut)
}

func (a *App) Run(host, port string) {
	listenAdress := fmt.Sprintf("%s:%s", host, port)

	fmt.Println("Started on: http://" + listenAdress)

	err := http.ListenAndServe(listenAdress, a.Router)

	utils.CheckError(err, "Error: http.ListenAndServe")
}

func main() {
	a := App{}

	// Load .env file
	err := godotenv.Load(".env")

	utils.CheckError(err, "Error: loading .env file")

	// Inits
	models.InitDB(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_NAME"))

	a.InitRoutes()

	a.Run(
		os.Getenv("HOST"),
		os.Getenv("PORT"))
}
