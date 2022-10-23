package models

import (
	"log"

	"github.com/Arcady1/go-rest-api/pkg/utils"
	"github.com/google/uuid"
)

func createAccount(amount float64) (string, error) {
	log.Println("models.createAccount", amount)

	accountId := uuid.New().String()

	_, err := DB.Query("INSERT INTO accounts(account_id, balance) VALUES ($1, $2);", accountId, amount)
	if err != nil {
		return "", err
	}

	return accountId, nil
}

func createUser(userId, accountId string) error {
	log.Println("models.createUser", userId, accountId)

	_, err := DB.Query("INSERT INTO users(user_id, account_id) VALUES ($1, $2);", userId, accountId)
	if err != nil {
		return err
	}

	return nil
}

func findUserById(user *users, userId string) error {
	log.Println("models.findUserById", userId)

	rows, err := DB.Query("SELECT * FROM users WHERE user_id = $1 LIMIT 1;", userId)
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(&user.UserId, &user.AccountId)
		if err != nil {
			return err
		}
	}
	rows.Close()

	return nil
}

func getUserBalanceByAccountId(accountId string) (float64, error) {
	log.Println("models.getUserBalanceByAccountId", accountId)

	var (
		account_id     string
		currentBalance float64
	)

	rows, err := DB.Query("SELECT * FROM accounts WHERE account_id = $1 LIMIT 1;", accountId)
	if err != nil {
		return 0, err
	}

	for rows.Next() {
		err = rows.Scan(&account_id, &currentBalance)
		if err != nil {
			return 0, err
		}
	}
	rows.Close()

	return currentBalance, nil
}

func updateUserAccountBalance(accountId string, balance float64) error {
	log.Println("models.updateUserAccountBalance", accountId, balance)

	balance, err := utils.PrepareAmountValue(balance)
	if err != nil {
		return err
	}

	_, err = DB.Query("UPDATE accounts SET balance = $1 WHERE account_id = $2;", balance, accountId)
	if err != nil {
		return err
	}

	return nil
}
