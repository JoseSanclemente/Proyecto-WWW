package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"Proyecto-WWW/back/models/consumer"
	"Proyecto-WWW/back/shared/gateway"

	"github.com/dchest/captcha"
)

func generateCaptcha(response http.ResponseWriter, request *http.Request) {
	captchaID := captcha.New()

	buf := bytes.NewBufferString("")
	encoder := base64.NewEncoder(base64.StdEncoding, buf)

	err := captcha.WriteImage(encoder, captchaID, captcha.StdWidth, captcha.StdHeight)
	if err != nil {
		fmt.Println("generateCaptcha_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}
	_ = encoder.Close()

	gateway.WriteJSON(response, 200, map[string]string{
		"captcha_id": captchaID,
		"captcha":    buf.String(),
	})
}

func loginConsumer(response http.ResponseWriter, request *http.Request) {
	params := map[string]string{}

	err := json.NewDecoder(request.Body).Decode(&params)
	if err != nil {
		fmt.Println("loginConsumer_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	if !captcha.VerifyString(params["captcha_id"], params["captcha"]) {
		gateway.WriteJSON(response, 400, map[string]string{"error": "wrong captcha"})
		return
	}

	ok, err := consumer.Login(params["id"], params["password"])
	if err != nil {
		fmt.Println("loginConsumer_2: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	if ok {
		gateway.WriteJSON(response, 200, map[string]string{"message": "ok"})
		return
	}

	gateway.WriteJSON(response, 400, map[string]string{"error": "invalid login parameters"})
}

func createConsumer(response http.ResponseWriter, request *http.Request) {
	input := consumer.Input{}

	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		fmt.Println("createConsumer_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	c := consumer.Consumer(input)
	newConsumer := &c

	err = newConsumer.Store()
	if err != nil {
		fmt.Println("createConsumer_2: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	gateway.WriteJSON(response, 200, map[string]string{"message": "ok"})
}

func updateConsumer(response http.ResponseWriter, request *http.Request) {
	params := map[string]string{}

	err := json.NewDecoder(request.Body).Decode(&params)
	if err != nil {
		fmt.Println("updateConsumer_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	c, err := consumer.Load(params["id"])
	if err != nil {
		fmt.Println("updateConsumer_2: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	if c == nil {
		gateway.WriteJSON(response, 400, map[string]string{"error": "user not found"})
		return
	}

	err = c.Update(params["password"], params["email"], params["deleted"])
	if err != nil {
		fmt.Println("updateConsumer_3: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	gateway.WriteJSON(response, 200, map[string]string{"message": "ok"})
}

func listConsumers(response http.ResponseWriter, request *http.Request) {
	consumers, err := consumer.List()
	if err != nil {
		fmt.Println("listConsumers_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	gateway.WriteJSON(response, 200, consumers)
}
