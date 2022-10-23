package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type JsonResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	ResponseErrWrongBodyFormat string = "Wrong body format"
	ResponseErrWrongData       string = "Wrong data"
	ResponseOK                 string = "ok"
)

func ResponseWriter(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	log.Println("utils.ResponseWriter")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := JsonResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
	}

	json.NewEncoder(w).Encode(response)
}

func CheckQuery(r *http.Request, value string, regExp string) error {
	log.Println("utils.CheckQuery")

	match, err := regexp.MatchString(regExp, value)

	if err != nil {
		return err
	}

	if !match {
		errString := fmt.Sprintf("Wrong param format. Got: %s, expected: %s", value, regExp)
		err := errors.New(errString)
		return err
	}

	return nil
}

func CheckError(err error, message string) {
	log.Println("utils.CheckError")

	if err != nil {
		log.Println(err)
		log.Fatalln(message)
	}
}

func PrepareAmountValue(amount float64) (float64, error) {
	log.Println("utils.PrepareAmountValue", amount)

	amountToStr := fmt.Sprintf("%.2f", amount)

	amount, err := strconv.ParseFloat(amountToStr, 64)
	if err != nil {
		return 0, err
	}

	return amount, nil
}
