package main

import "os"

func main() {
	a := App{}

	a.InitDB(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.InitRoutes()

	a.Run(
		os.Getenv("HOST"),
		os.Getenv("PORT"))
}
