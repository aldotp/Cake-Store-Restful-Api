package helper

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, httpCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(response)
}

func ResponseError(w http.ResponseWriter, httpCode int, err error) {
	if httpCode == http.StatusInternalServerError {
		log.Println("Terjadi error =", err.Error())
		err = errors.New("internal server error")
	}

	ResponseJSON(w, httpCode, map[string]string{"message": err.Error()})
}
