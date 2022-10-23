package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Arcady1/go-rest-api/pkg/models"
	"github.com/Arcady1/go-rest-api/pkg/utils"
)

type balance struct {
	UserId string `json:"userId"`
}

func GetUserBalance(w http.ResponseWriter, r *http.Request) {
	log.Println("handlers.GetUserBalance")

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, http.StatusInternalServerError, utils.ResponseErrWrongBodyFormat, nil)
		return
	}

	// Try to save the user ID in a variable
	var userBalance balance

	err = json.Unmarshal(body, &userBalance)
	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, http.StatusBadRequest, utils.ResponseErrWrongBodyFormat, nil)
		return
	}

	// Validate userId
	err = utils.CheckQuery(r, userBalance.UserId, bodyPatterns["UserId"])
	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, http.StatusBadRequest, utils.ResponseErrWrongData, nil)
		return
	}

	// Close body
	r.Body.Close()

	// Get the user balance
	var (
		data       interface{}
		statusCode int
	)

	data, err, statusCode = models.GetAccountBalance(userBalance.UserId)
	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, statusCode, ResponseErrRefillUserAccount, nil)
		return
	}

	// Send a response
	utils.ResponseWriter(w, http.StatusOK, utils.ResponseSuccess, data)
}
