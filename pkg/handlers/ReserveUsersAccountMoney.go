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

type reserve struct {
	UserId    string  `json:"userId"`
	ServiceId string  `json:"serviceId"`
	OrderId   string  `json:"orderId"`
	Cost      float64 `json:"cost"`
}

const (
	ResponseErrReserveUsersAccountMoney string = "Error: reserving money"
)

func ReserveUsersAccountMoney(w http.ResponseWriter, r *http.Request) {
	log.Println("handlers.ReserveUsersAccountMoney")

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, http.StatusInternalServerError, utils.ResponseErrWrongBodyFormat, nil)
		return
	}

	// Try to save the reserve data in a variable
	var moneyReserve reserve

	err = json.Unmarshal(body, &moneyReserve)
	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, http.StatusBadRequest, utils.ResponseErrWrongBodyFormat, nil)
		return
	}

	// Validate userId
	err = utils.CheckQuery(r, moneyReserve.UserId, bodyPatterns["UserId"])
	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, http.StatusBadRequest, utils.ResponseErrWrongData, nil)
		return
	}

	// Validate ServiceId
	err = utils.CheckQuery(r, moneyReserve.ServiceId, bodyPatterns["ServiceId"])
	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, http.StatusBadRequest, utils.ResponseErrWrongData, nil)
		return
	}

	// Validate OrderId
	err = utils.CheckQuery(r, moneyReserve.OrderId, bodyPatterns["OrderId"])
	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, http.StatusBadRequest, utils.ResponseErrWrongData, nil)
		return
	}

	// Validate Cost
	err = utils.CheckQuery(r, fmt.Sprintf("%v", moneyReserve.Cost), bodyPatterns["Amount"])
	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, http.StatusBadRequest, utils.ResponseErrWrongData, nil)
		return
	}

	// Prepare cost value
	moneyReserve.Cost, err = utils.PrepareAmountValue(moneyReserve.Cost)
	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, http.StatusInternalServerError, ResponseErrReserveUsersAccountMoney, nil)
		return
	}

	// Close body
	r.Body.Close()

	// Reserve money
	var statusCode int

	err, statusCode = models.ReserveUsersAccountMoney(moneyReserve.UserId, moneyReserve.ServiceId, moneyReserve.OrderId, moneyReserve.Cost)
	if err != nil {
		log.Println(err)
		errMessage := ResponseErrReserveUsersAccountMoney

		if statusCode != http.StatusInternalServerError {
			errMessage = err.Error()
		}

		utils.ResponseWriter(w, statusCode, errMessage, nil)
		return
	}

	// Send a response
	utils.ResponseWriter(w, http.StatusOK, utils.ResponseSuccess, nil)
}
