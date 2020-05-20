package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"univalle/www/models/transformer"
	"univalle/www/shared/gateway"
)

func createTransformer(response http.ResponseWriter, request *http.Request) {
	newTransformer := &transformer.Transformer{}

	err := json.NewDecoder(request.Body).Decode(newTransformer)
	if err != nil {
		fmt.Println("createTransformer_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	id, err := newTransformer.Store()
	if err != nil {
		fmt.Println("createTransformer_2: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	gateway.WriteJSON(response, 200, map[string]string{"id": id})
}

func updateTransformer(response http.ResponseWriter, request *http.Request) {
	params := map[string]string{}

	err := json.NewDecoder(request.Body).Decode(&params)
	if err != nil {
		fmt.Println("updateTransformer_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	t, err := transformer.Load(params["id"])
	if err != nil {
		fmt.Println("updateTransformer_2: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	if t == nil {
		gateway.WriteJSON(response, 400, map[string]string{"error": "user not found"})
		return
	}

	err = t.Update(params["name"], params["longitude"], params["latitude"], params["deleted"])
	if err != nil {
		fmt.Println("updateTransformer_3: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	gateway.WriteJSON(response, 200, map[string]string{"message": "ok"})
}

func listTransformers(response http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("substation_id")
	if id == "" {
		response.Header().Set("Access-Control-Allow-Origin", "*")
		gateway.WriteJSON(response, 400, map[string]string{"message": "missing substation id"})
		return
	}

	transformers, err := transformer.List(id)
	if err != nil {
		fmt.Println("listTransformers_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	gateway.WriteJSON(response, 200, transformers)
}
