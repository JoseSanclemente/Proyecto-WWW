package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"Proyecto-WWW/back/models/contract"
	"Proyecto-WWW/back/shared/gateway"
)

func createContract(response http.ResponseWriter, request *http.Request) {
	newContract := &contract.Contract{}

	err := json.NewDecoder(request.Body).Decode(newContract)
	if err != nil {
		fmt.Println("createContract_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	id, err := newContract.Store()
	if err != nil {
		fmt.Println("createContract_2: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	gateway.WriteJSON(response, 200, map[string]string{"id": id})
}

func updateContract(response http.ResponseWriter, request *http.Request) {
	params := map[string]string{}

	err := json.NewDecoder(request.Body).Decode(&params)
	if err != nil {
		fmt.Println("updateContract_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	c, err := contract.Load(params["id"])
	if err != nil {
		fmt.Println("updateContract_2: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	if c == nil {
		gateway.WriteJSON(response, 400, map[string]string{"error": "contract not found"})
		return
	}

	err = c.Update(params["notification_type"], params["deleted"])
	if err != nil {
		fmt.Println("updateContract_3: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	gateway.WriteJSON(response, 200, map[string]string{"message": "ok"})
}

func listContracts(response http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("consumer_id")
	if id == "" {
		response.Header().Set("Access-Control-Allow-Origin", "*")
		gateway.WriteJSON(response, 400, map[string]string{"message": "missing consumer id"})
		return
	}

	contracts, err := contract.List(id)
	if err != nil {
		fmt.Println("listConsumers_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	gateway.WriteJSON(response, 200, contracts)
}
