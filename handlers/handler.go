package handlers

import (
	"net/http"

	"github.com/FadyGamilM/Rss-Aggregator/helpers"
)

func Test_server_health(w http.ResponseWriter, r *http.Request) {
	helpers.JsonResponse(w, 200, "hello")
}

func Test_error_handler(w http.ResponseWriter, r *http.Request) {
	helpers.ErrorResponse(w, 500, "internal server error")
}
