package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Arcady1/go-rest-api/pkg/models"
)

func IncreaseUserBalance(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		err = errors.New("Error: decode request body")
		log.Fatalln(err)
	}

	var balance models.Balance
	err = json.Unmarshal(body, &balance)

	if err != nil {
		err = errors.New("Error: wrong body format")
		log.Fatalln(err)
	}

	// Increase the user balance
	// TODO

	// Send a response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Done")
}
