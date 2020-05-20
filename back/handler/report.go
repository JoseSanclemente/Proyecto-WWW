package handler

import (
	"net/http"
)

func report(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.WriteHeader(200)
}
