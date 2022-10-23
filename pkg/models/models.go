package models

import (
	"database/sql"
	"fmt"

	"github.com/Arcady1/go-rest-api/pkg/utils"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(user, password, dbname string) {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", dbInfo)

	utils.CheckError(err, "Error: connecting to the Database")
}
