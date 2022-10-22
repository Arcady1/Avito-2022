package models

import (
	
)

db := 

func ConnectToDB() (*sql.DB, error) {
	USER := os.Getenv("DB_USER")
	PASSWORD := os.Getenv("DB_PASSWORD")
	DBNAME := os.Getenv("DB_NAME")
	SSLMODE := os.Getenv("DB_SSLMODE")

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", USER, PASSWORD, DBNAME, SSLMODE)
	database, err := sql.Open("postgres", dbinfo)

	if err != nil {
		return nil, errors.New("Error: connect to the Database")
	} else {
		db = database
	}

	return db, nil
}
