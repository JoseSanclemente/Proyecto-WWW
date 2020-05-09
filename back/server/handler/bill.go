package handler

import (
	"Proyecto-WWW/back/models/bill"
	"Proyecto-WWW/back/shared/gateway"
	"encoding/json"
	"fmt"
	"net/http"
)

func registerPayment(response http.ResponseWriter, request *http.Request) {
	params := map[string]string{}

	err := json.NewDecoder(request.Body).Decode(&params)
	if err != nil {
		fmt.Println("updateBill_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	b, err := bill.Load(params["id"])
	if err != nil {
		fmt.Println("updateBill_2: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	if b == nil {
		gateway.WriteJSON(response, 400, map[string]string{"error": "bill not found"})
		return
	}

	err = b.RegisterPayment()
	if err != nil {
		fmt.Println("updateBill_3: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	gateway.WriteJSON(response, 200, map[string]string{"message": "ok"})
}

func createBill(response http.ResponseWriter, request *http.Request) {
	newBill := &bill.Bill{}

	err := json.NewDecoder(request.Body).Decode(newBill)
	if err != nil {
		fmt.Println("createBill_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	err = newBill.Store()
	if err != nil {
		fmt.Println("createBill_2: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	gateway.WriteJSON(response, 200, map[string]string{"message": "ok"})
}
