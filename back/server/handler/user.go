package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"Proyecto-WWW/back/models/user"
	"Proyecto-WWW/back/shared/gateway"
)

func loginUser(response http.ResponseWriter, request *http.Request) {
	params := map[string]string{}

	err := json.NewDecoder(request.Body).Decode(&params)
	if err != nil {
		fmt.Println("loginUser_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	userType, err := user.Login(params["email"], params["password"])
	if err != nil {
		fmt.Println("loginUser_2: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	if userType == "" {
		gateway.WriteJSON(response, 400, map[string]string{"error": "invalid login parameters"})
		return
	}

	gateway.WriteJSON(response, 200, map[string]string{"type": userType})
}

func createUser(response http.ResponseWriter, request *http.Request) {
	newUser := &user.User{}

	err := json.NewDecoder(request.Body).Decode(newUser)
	if err != nil {
		fmt.Println("createUser_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	err = newUser.Store()
	if err != nil {
		fmt.Println("createUser_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	gateway.WriteJSON(response, 200, map[string]string{"message": "ok"})
}

func updateUser(response http.ResponseWriter, request *http.Request) {
	params := map[string]string{}

	err := json.NewDecoder(request.Body).Decode(&params)
	if err != nil {
		fmt.Println("updateUser_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	u, err := user.Load(params["email"])
	if err != nil {
		fmt.Println("updateUser_2: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	if u == nil {
		gateway.WriteJSON(response, 400, map[string]string{"error": "user not found"})
		return
	}

	err = u.Update(params["password"], params["type"], params["deleted"])
	if err != nil {
		fmt.Println("updateUser_3: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	gateway.WriteJSON(response, 200, map[string]string{"message": "ok"})
}

func listUsers(response http.ResponseWriter, request *http.Request) {
	users, err := user.List()
	if err != nil {
		fmt.Println("updateUser_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	gateway.WriteJSON(response, 200, users)
}
