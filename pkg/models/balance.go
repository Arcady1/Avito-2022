package models

import "log"

type accounts struct {
	UserId string  `json:"userId"`
	Amount float64 `json:"amount"`
}

func RefillUserAccount(userId string, amount float64) (string, error) {
	log.Println("models.RefillUserAccount", userId, amount)

	var user users

	// Find a user with 'userId'
	err := findUserById(&user, userId)
	if err != nil {
		return "", err
	}

	// Create an account and a user if the user doesn't exist
	if (users{}) == user {
		// Create an account
		var accountId string

		accountId, err = createAccount(amount)
		if err != nil {
			return "", err
		}

		// Create a user
		err = createUser(userId, accountId)
		if err != nil {
			return "", err
		}
	} else {
		var accountBalance, newBalance float64

		// Get a user account ID
		accountId := user.AccountId

		// Get a current user account balance
		accountBalance, err = getUserBalanceByAccountId(accountId)
		if err != nil {
			return "", err
		}

		newBalance = accountBalance + amount

		// Update a user account balance
		err = updateUserAccountBalance(accountId, newBalance)
		if err != nil {
			return "", err
		}
	}

	return "okay", nil
}
