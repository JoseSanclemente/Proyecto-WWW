package handler

import (
	"net/http"
)

func report(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
}
