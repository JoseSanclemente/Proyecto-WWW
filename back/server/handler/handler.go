package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	getEndpoints = map[string]func(response http.ResponseWriter, request *http.Request){
		"/employee/list":          listUsers,
		"/substation/list":        listSubstations,
		"/transformer/list":       listTransformers,
		"/consumer/captcha":       generateCaptcha,
		"/consumer/list":          listConsumers,
		"/consumer/contract/list": listContracts,
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
	}

	updateEndpoints = map[string]func(response http.ResponseWriter, request *http.Request){
		"/employee":          updateUser,
		"/transformer":       updateTransformer,
		"/substation":        updateSubstation,
		"/consumer":          updateConsumer,
		"/consumer/contract": updateContract,
	}

	endpointsLists = map[string]map[string]func(response http.ResponseWriter, request *http.Request){
		"GET":    getEndpoints,
		"POST":   postEndpoints,
		"UPDATE": updateEndpoints,
	}
)

// NewRouter with handles
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	for method, endpoints := range endpointsLists {
		for path, endpoint := range endpoints {
			router.HandleFunc(path, corsHandler).Methods("OPTIONS")
			router.HandleFunc(path, endpoint).Methods(method)
		}
	}

	return router
}
