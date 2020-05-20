package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	getEndpoints    map[string]func(response http.ResponseWriter, request *http.Request)
	postEndpoints   map[string]func(response http.ResponseWriter, request *http.Request)
	updateEndpoints map[string]func(response http.ResponseWriter, request *http.Request)
	endpointsLists  map[string]map[string]func(response http.ResponseWriter, request *http.Request)
)

func init() {
	getEndpoints = map[string]func(response http.ResponseWriter, request *http.Request){
		"/":                       rootHandler,
		"/employee/list":          listUsers,
		"/substation/list":        listSubstations,
		"/transformer/list":       listTransformers,
		"/consumer/captcha":       generateCaptcha,
		"/consumer/list":          listConsumers,
		"/consumer/contract/list": listContracts,
		"/consumer/contract/pdf":  getPDF,
		"/dummy_report":           report,
	}

	postEndpoints = map[string]func(response http.ResponseWriter, request *http.Request){
		"/employee/login":    loginUser,
		"/employee":          createUser,
		"/substations":       createSubstation,
		"/transformer":       createTransformer,
		"/consumer/login":    loginConsumer,
		"/consumer":          createConsumer,
		"/consumer/contract": createContract,
		"/bill":              createBill,
	}

	updateEndpoints = map[string]func(response http.ResponseWriter, request *http.Request){
		"/employee":              updateUser,
		"/transformer":           updateTransformer,
		"/substation":            updateSubstation,
		"/consumer":              updateConsumer,
		"/consumer/contract":     updateContract,
		"/bill/payment/operator": registerOperatorPayment,
		"/bill/payment/bank":     registerBankPayments,
	}

	endpointsLists = map[string]map[string]func(response http.ResponseWriter, request *http.Request){
		"GET":  getEndpoints,
		"POST": postEndpoints,
		"PUT":  updateEndpoints,
	}
}

func rootHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)

	var out string
	for method, endpoints := range endpointsLists {
		for path, _ := range endpoints {
			out += method + " " + path + "\n"
		}
		out += "\n"
	}

	response.Write([]byte(out))
}

// NewRouter with handles
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)

	for method, endpoints := range endpointsLists {
		for path, endpoint := range endpoints {
			router.HandleFunc(path, corsHandler).Methods("OPTIONS")
			router.HandleFunc(path, endpoint).Methods(method)
		}
	}

	return router
}
