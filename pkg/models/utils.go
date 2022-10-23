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

func findOrder(order *orders, accountId, serviceId, orderId string) error {
	log.Println("models.findOrder", accountId, serviceId, orderId)

	rows, err := DB.Query("SELECT * FROM orders WHERE (account_id = $1) AND (order_id = $2) AND (service_id = $3) LIMIT 1;", accountId, orderId, serviceId)
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(&order.Id, &order.OrderId, &order.AccountId, &order.ServiceId, &order.Cost, &order.Status)
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

func createdNewOrder(accountId, serviceId, orderId string, cost float64) error {
	log.Println("models.createdNewOrder", accountId, serviceId, orderId, cost)

	id := uuid.New().String()

	_, err := DB.Query("INSERT INTO orders(id, order_id, account_id, service_id, cost, status) VALUES ($1, $2, $3, $4, $5, 'reserved');", id, orderId, accountId, serviceId, cost)
	if err != nil {
		return err
	}

	return nil
}

func reserveMoney(accountId, serviceId, orderId string, cost, accountBalance float64) error {
	log.Println("models.reserveMoney", accountId, serviceId, orderId, cost)

	// Reserve money
	err := createdNewOrder(accountId, serviceId, orderId, cost)
	if err != nil {
		return err
	}

	// Write-off money from the user's account if there is no error on reserving money
	newBalance := accountBalance - cost

	err = updateUserAccountBalance(accountId, newBalance)
	if err != nil {
		return err
	}

	return nil
}

func updateOrder(accountId, serviceId, orderId, status string, amount float64) error {
	log.Println("models.updateOrder", accountId, serviceId, orderId, status)

	// Update order status 'canceled'
	if status == "canceled" {
		_, err := DB.Query("UPDATE orders SET status = 'canceled' WHERE (account_id = $1) AND (order_id = $2) AND (service_id = $3);", accountId, orderId, serviceId)
		if err != nil {
			return err
		}
	} else if status == "succeed" {
		// Update order status 'succeed'
		_, err := DB.Query("UPDATE orders SET status = 'succeed', cost = $1 WHERE (account_id = $2) AND (order_id = $3) AND (service_id = $4);", amount, accountId, orderId, serviceId)
		if err != nil {
			return err
		}
	}

	return nil
}

func updateMoneyReserve(accountId, serviceId, orderId, status string, orderCost, amount float64) error {
	log.Println("models.updateMoneyReserve", accountId, serviceId, orderId, orderCost, amount)

	// Update the order
	err := updateOrder(accountId, serviceId, orderId, status, amount)
	if err != nil {
		return err
	}

	// Refund the user if it is impossible to write off the money OR 'amount < orderCost'
	if status == "succeed" {
		var (
			accountBalance float64
			newBalance     float64
		)

		// Get the user's account balance
		accountBalance, err = getUserBalanceByAccountId(accountId)
		if err != nil {
			return err
		}

		// Update the user's account balance. 'orderCost - amount' != 0 when the 'amount < orderCost'
		newBalance = accountBalance + (orderCost - amount)

		err = updateUserAccountBalance(accountId, newBalance)
		if err != nil {
			return err
		}
	}

	return nil
}
