package models

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

type accounts struct {
	UserId string  `json:"userId"`
	Amount float64 `json:"amount"`
}

func RefillUserAccount(userId string, amount float64) error {
	log.Println("models.RefillUserAccount", userId, amount)

	var user users

	// Find a user with 'userId'
	err := findUserById(&user, userId)
	if err != nil {
		return err
	}

	// Create an account and a user if the user doesn't exist
	if (users{}) == user {
		// Create an account
		var accountId string

		accountId, err = createAccount(amount)
		if err != nil {
			return err
		}

		// Create a user
		err = createUser(userId, accountId)
		if err != nil {
			return err
		}
	} else {
		var accountBalance, newBalance float64

		// Get a user account ID
		accountId := user.AccountId

		// Get a current user account balance
		accountBalance, err = getUserBalanceByAccountId(accountId)
		if err != nil {
			return err
		}

		newBalance = accountBalance + amount

		// Update a user account balance
		err = updateUserAccountBalance(accountId, newBalance)
		if err != nil {
			return err
		}
	}

	// Success
	return nil
}

func GetAccountBalance(userId string) (interface{}, error, int) {
	log.Println("models.GetAccountBalance", userId)

	// Find the user by userId
	var user users

	// Find a user with 'userId'
	err := findUserById(&user, userId)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	// If the user doesn't exist, return an error
	if (users{}) == user {
		errMessage := fmt.Sprintf("The user with ID '%s' is not found", userId)
		err = errors.New(errMessage)
		return nil, err, http.StatusNotFound
	}

	// Get the account ID
	accountId := user.AccountId

	// Get the balance by account ID
	var accountBalance float64

	accountBalance, err = getUserBalanceByAccountId(accountId)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	// Success
	response := map[string]float64{
		"balance": accountBalance,
	}
	return response, nil, http.StatusOK
}
