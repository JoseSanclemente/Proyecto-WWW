package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"Proyecto-WWW/back/server/handler"
	"Proyecto-WWW/back/storage"
)

func main() {
	var err error
	if len(os.Args) > 1 && os.Args[1] == "persist" {
		err = storage.Connect()
	} else {
		err = storage.ResetAndConnect()
	}

	if err != nil {
		panic(err)
	}
	defer storage.DB.Close()

	server := &http.Server{
		Addr:           ":8080",
		Handler:        handler.NewRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Starting server")
	log.Fatal(server.ListenAndServe())
}
