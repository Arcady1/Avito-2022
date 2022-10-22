package models

type Balance struct {
	UserId string  `json:"userId"`
	Amount float32 `json:"amount"`
}
