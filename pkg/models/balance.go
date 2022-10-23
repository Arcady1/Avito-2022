package models

type Balance struct {
	UserId string  `json:"userId"`
	Amount float32 `json:"amount"`
}

var BalancePatterns = map[string]string{
	"UserId": ".+",
	"Amount": "([1-9][0-9]*(\\.[0-9]{1,2})?)|(0\\.(0[1-9])|([1-9][0-9]?))",
}
