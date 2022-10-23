package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Arcady1/go-rest-api/pkg/models"
	"github.com/Arcady1/go-rest-api/pkg/utils"
)

type refill struct {
	UserId string  `json:"userId"`
	Amount float64 `json:"amount"`
}

var refillPatterns = map[string]string{
	"UserId": "^(.+)$",
	"Amount": "^(([1-9][0-9]*(\\.[0-9]{1,2})?)|(0\\.((0[1-9])|([1-9][0-9]?))))$",
}

const (
	ResponseErrRefillUserAccount string = "Error on refill user account"
)

func RefillUserAccount(w http.ResponseWriter, r *http.Request) {
	log.Println("handlers.RefillUserAccount")

	var successResponse string

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, 500, utils.ResponseErrWrongBodyFormat, nil)
		return
	}

	// Try to save the amount in a variable
	var refillAccount refill
	err = json.Unmarshal(body, &refillAccount)

	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, 400, utils.ResponseErrWrongBodyFormat, nil)
		return
	}

	// Validate userId
	err = utils.CheckQuery(r, refillAccount.UserId, refillPatterns["UserId"])
	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, 400, utils.ResponseErrWrongData, nil)
		return
	}

	// Validate amount
	err = utils.CheckQuery(r, fmt.Sprintf("%v", refillAccount.Amount), refillPatterns["Amount"])
	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, 400, utils.ResponseErrWrongData, nil)
		return
	}

	// Close body
	r.Body.Close()

	// Prepare amount value
	refillAccount.Amount, err = utils.PrepareAmountValue(refillAccount.Amount)
	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, 500, ResponseErrRefillUserAccount, nil)
		return
	}

	// Refill the user account
	successResponse, err = models.RefillUserAccount(refillAccount.UserId, refillAccount.Amount)
	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, 500, ResponseErrRefillUserAccount, nil)
		return
	}

	// Send a response
	utils.ResponseWriter(w, http.StatusOK, successResponse, nil)
}
