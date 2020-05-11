package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"Proyecto-WWW/back/models/bill"
	"Proyecto-WWW/back/models/contract"
	"Proyecto-WWW/back/models/reading"
	"Proyecto-WWW/back/server/handler"
	"Proyecto-WWW/back/storage"
)

func createReadingsAndBills() {
	contracts, err := contract.ListActiveContractsIDs()
	if err != nil {
		fmt.Println("createReadingsAndBills_1: ", err.Error())
		return
	}

	for _, contract := range contracts {
		active, err := bill.IsContractActive(contract.ID)
		if err != nil {
			fmt.Println("createReadingsAndBills_2: ", err.Error())
			continue
		}

		if !active {
			continue
		}

		r := &reading.Reading{
			ContractID: contract.ID,
			Value:      rand.Intn(4),
			Date:       int(time.Now().Unix()),
		}

		err = r.Store()
		if err != nil {
			fmt.Println("createReadingsAndBills_3: ", err.Error())
			continue
		}
	}
}

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

	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("rutina de readings")
		}
	}()

	fmt.Println("Starting server")
	log.Fatal(server.ListenAndServe())
}
