package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
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
	match, err := regexp.MatchString(regExp, value)

	if (!match) || (err != nil) {
		log.Println(err)
		errString := fmt.Sprintf("Wrong param format. Got: %s, expected: %s", value, regExp)
		err := errors.New(errString)
		return err
	}

	return nil
}
