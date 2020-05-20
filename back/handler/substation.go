package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"univalle/www/models/substation"
	"univalle/www/shared/gateway"
)

func createSubstation(response http.ResponseWriter, request *http.Request) {
	newSubstation := &substation.Substation{}

	err := json.NewDecoder(request.Body).Decode(newSubstation)
	if err != nil {
		fmt.Println("createSubstation_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	id, err := newSubstation.Store()
	if err != nil {
		fmt.Println("createSubstation_2: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	gateway.WriteJSON(response, 200, map[string]string{"id": id})
}

func updateSubstation(response http.ResponseWriter, request *http.Request) {
	params := map[string]string{}

	err := json.NewDecoder(request.Body).Decode(&params)
	if err != nil {
		fmt.Println("updateSubstation_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	s, err := substation.Load(params["id"])
	if err != nil {
		fmt.Println("updateSubstation_2: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	if s == nil {
		gateway.WriteJSON(response, 400, map[string]string{"error": "user not found"})
		return
	}

	err = s.Update(params["name"], params["longitude"], params["latitude"], params["deleted"])
	if err != nil {
		fmt.Println("updateSubstation_3: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	gateway.WriteJSON(response, 200, map[string]string{"message": "ok"})
}

func listSubstations(response http.ResponseWriter, request *http.Request) {
	substations, err := substation.List()
	if err != nil {
		fmt.Println("listSubstation_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	gateway.WriteJSON(response, 200, substations)
}
