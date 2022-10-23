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
	err = reserveMoney(accountId, serviceId, orderId, cost, accountBalance)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	// Success
	return nil, http.StatusOK
}

func AcceptReservedMoney(userId, serviceId, orderId string, amount float64) (error, int) {
	log.Println("models.AcceptReservetMoney", userId, serviceId, orderId, amount)

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

	// Find an order by accountId, serviceId and orderId
	var order orders

	err = findOrder(&order, accountId, serviceId, orderId)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	// If the order doesn't exist, return an error
	if (orders{}) == order {
		err = errors.New("The order is not found")
		return err, http.StatusNotFound
	}

	// If the order status is not 'reserved', return an error
	if order.Status != "reserved" {
		errorMessage := fmt.Sprintf("The order is not reserved. The order status is '%s'", order.Status)
		err = errors.New(errorMessage)
		return err, http.StatusNotAcceptable
	}

	// Update money reserve status
	var reserveStatus string

	if amount > order.Cost {
		reserveStatus = "canceled"
	} else {
		// Accept reserved money
		reserveStatus = "succeed"
	}

	// Update money reserve
	err = updateMoneyReserve(accountId, serviceId, orderId, reserveStatus, order.Cost, amount)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	// Return
	if reserveStatus == "canceled" {
		err = errors.New("The 'amount' is greater than the reserved amount of money. The order status is 'canceled'. Refund the user")
		return err, http.StatusCreated
	}

	return nil, http.StatusCreated
}
