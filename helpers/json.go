package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func ErrorResponse(w http.ResponseWriter, statusCode int, errMsg string) {
	if statusCode > 499 {
		log.Printf("[X][LOGGING] respond to 5xx error : %v \n", errMsg)
	}

	JsonResponse(w, statusCode, errorResponse{Error: errMsg})
}

func JsonResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	// marshel the payload from go data sturcture to json
	resp, err := json.Marshal(payload)
	if err != nil {
		log.Printf("[X][LOGGING] error while converting data to json response : %v \n", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resp)
}
