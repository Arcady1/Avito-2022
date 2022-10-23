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

func IncreaseUserBalance(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, 500, utils.ResponseErrWrongBodyFormat, nil)
		return
	}

	// Try to save balance in a variable
	var balance models.Balance
	err = json.Unmarshal(body, &balance)

	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, 400, utils.ResponseErrWrongBodyFormat, nil)
		return
	}

	// Validate userId
	err = utils.CheckQuery(r, balance.UserId, models.BalancePatterns["UserId"])
	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, 400, utils.ResponseErrWrongData, nil)
		return
	}

	// Validate amount
	err = utils.CheckQuery(r, fmt.Sprintf("%v", balance.Amount), models.BalancePatterns["Amount"])
	if err != nil {
		log.Println(err)
		utils.ResponseWriter(w, 400, utils.ResponseErrWrongData, nil)
		return
	}

	// Increase the user balance
	// TODO

	// Send a response
	utils.ResponseWriter(w, http.StatusOK, "TODO", nil)
}
