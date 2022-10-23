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

type acceptReserved struct {
	UserId    string  `json:"userId"`
	ServiceId string  `json:"serviceId"`
	OrderId   string  `json:"orderId"`
	Amount    float64 `json:"amount"`
}

const (
	ResponseErrAcceptReservedMoney string = "Error: accepting reserved money"
)

func AcceptReservedMoney(w http.ResponseWriter, r *http.Request) {
	log.Println("handlers.AcceptReservedMoney")

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, http.StatusInternalServerError, utils.ResponseErrWrongBodyFormat, nil)
		return
	}

	// Try to save the reserve data in a variable
	var moneyAccept acceptReserved

	err = json.Unmarshal(body, &moneyAccept)
	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, http.StatusBadRequest, utils.ResponseErrWrongBodyFormat, nil)
		return
	}

	// Validate userId
	err = utils.CheckQuery(r, moneyAccept.UserId, bodyPatterns["UserId"])
	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, http.StatusBadRequest, utils.ResponseErrWrongData, nil)
		return
	}

	// Validate ServiceId
	err = utils.CheckQuery(r, moneyAccept.ServiceId, bodyPatterns["ServiceId"])
	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, http.StatusBadRequest, utils.ResponseErrWrongData, nil)
		return
	}

	// Validate OrderId
	err = utils.CheckQuery(r, moneyAccept.OrderId, bodyPatterns["OrderId"])
	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, http.StatusBadRequest, utils.ResponseErrWrongData, nil)
		return
	}

	// Validate Cost
	err = utils.CheckQuery(r, fmt.Sprintf("%v", moneyAccept.Amount), bodyPatterns["Amount"])
	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, http.StatusBadRequest, utils.ResponseErrWrongData, nil)
		return
	}

	// Close body
	r.Body.Close()

	// Accept reserved money or refund the user
	var statusCode int

	err, statusCode = models.AcceptReservedMoney(moneyAccept.UserId, moneyAccept.ServiceId, moneyAccept.OrderId, moneyAccept.Amount)
	if err != nil {
		log.Println(err)
		errMessage := ResponseErrAcceptReservedMoney

		if statusCode != http.StatusInternalServerError {
			errMessage = err.Error()
		}

		utils.ResponseWriter(w, statusCode, errMessage, nil)
		return
	}

	// Send a response
	utils.ResponseWriter(w, http.StatusCreated, utils.ResponseSuccess, nil)
}
