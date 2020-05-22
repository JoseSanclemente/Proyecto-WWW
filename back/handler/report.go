package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"Proyecto-WWW/back/models/report"
	"Proyecto-WWW/back/shared/gateway"
)

func monthly(response http.ResponseWriter, request *http.Request) {
	params := map[string]int{}

	err := json.NewDecoder(request.Body).Decode(&params)
	if err != nil {
		fmt.Println("monthly_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	result, err := report.GetMonthlyConsumption(params["start_year"], params["start_month"], params["end_year"], params["end_month"])
	if err != nil {
		fmt.Println("monthly_2: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	gateway.WriteJSON(response, 200, result)
}

func yearly(response http.ResponseWriter, request *http.Request) {
	params := map[string]int{}

	err := json.NewDecoder(request.Body).Decode(&params)
	if err != nil {
		fmt.Println("yearly_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	result, err := report.GetYearlyConsumption(params["start_year"], params["end_year"])
	if err != nil {
		fmt.Println("yearly_2: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	gateway.WriteJSON(response, 200, result)
}

func top(response http.ResponseWriter, request *http.Request) {
	result, err := report.GetTopConsumers()
	if err != nil {
		fmt.Println("top_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	gateway.WriteJSON(response, 200, result)
}
