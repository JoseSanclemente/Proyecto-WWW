package gateway

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteInternalServerError(r http.ResponseWriter) {
	r.WriteHeader(http.StatusInternalServerError)
	_, err := r.Write([]byte(`{"error":"something went wrong"}`))
	if err != nil {
		fmt.Println("writeInternalServerError: ", err.Error())
	}
}

func WriteJSON(r http.ResponseWriter, code int, body interface{}) {
	r.Header().Set("Content-Type", "application/json")
	r.Header().Set("Access-Control-Allow-Origin", "*")

	data, err := json.Marshal(body)
	if err != nil {
		WriteInternalServerError(r)
		return
	}

	r.WriteHeader(code)
	_, err = r.Write(data)
	if err != nil {
		fmt.Println("WriteJSON: ", err.Error())
		WriteInternalServerError(r)
		return
	}
}
