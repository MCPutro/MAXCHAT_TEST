package utils

import (
	"encoding/json"
	"net/http"
)

type ResponsePayload struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func WriteToResponseBody(err error, message string, w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)

	var status int
	var pesan string
	if err != nil {
		status = http.StatusInternalServerError
		pesan = message

	} else {
		status = http.StatusOK
		pesan = "success"
	}

	r := ResponsePayload{
		Status:  status,
		Message: pesan,
		Data:    response,
	}

	encoder.Encode(r)
}
