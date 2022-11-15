package helper

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, httpCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(response)
}

func ResponseError(w http.ResponseWriter, httpCode int, message string) {
	ResponseJSON(w, httpCode, map[string]string{"message": message})
}
