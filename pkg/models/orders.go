package models

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

type orders struct {
	OrderId   string  `json:"orderId"`
	AccountId string  `json:"accountId"`
	ServiceId string  `json:"serviceId"`
	Cost      float64 `json:"cost"`
	Status    string  `json:"status"`
}

func ReserveUsersAccountMoney(userId, serviceId, orderId string, cost float64) (error, int) {
	log.Println("models.ReserveUsersAccountMoney", userId, serviceId, orderId, cost)

	// Find the user by userId
	var user users

	// Find a user with 'userId'
	err := findUserById(&user, userId)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	// If the user doesn't exist, return an error
	if (users{}) == user {
		errMessage := fmt.Sprintf("The user with ID '%s' is not found", userId)
		err = errors.New(errMessage)
		return err, http.StatusNotFound
	}

	// Get the account ID
	accountId := user.AccountId

	// Get the balance by account ID
	var accountBalance float64

	accountBalance, err = getUserBalanceByAccountId(accountId)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	// If the user's balance less than the cost
	if accountBalance < cost {
		err = errors.New("The 'cost' is greater than the user's balance")
		return err, http.StatusNotAcceptable
	}

	// Reserve money
	err = ReserveMoney(accountId, serviceId, orderId, cost, accountBalance)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	// Success
	return nil, http.StatusOK
}
