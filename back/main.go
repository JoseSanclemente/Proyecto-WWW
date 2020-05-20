package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"univalle/www/handler"
	"univalle/www/models/bill"
	"univalle/www/models/contract"
	"univalle/www/models/reading"
	reports "univalle/www/models/report"
	"univalle/www/models/report/email"
	"univalle/www/storage"
)

func createReadings(contractID string) {
	active, err := bill.IsContractActive(contractID)
	if err != nil {
		fmt.Println("createReadings_1: ", err.Error())
		return
	}

	if !active {
		return
	}

	r := &reading.Reading{
		ContractID: contractID,
		Value:      rand.Intn(4),
		Date:       int(time.Now().Unix()),
	}

	err = r.Store()
	if err != nil {
		fmt.Println("createReadings_2: ", err.Error())
		return
	}
}

func createBills(contract *contract.Contract) {
	creationDate := time.Now().UTC()
	expirationDate := creationDate.Add(time.Hour * 24 * 15).Unix()

	value, err := reading.LoadTotal(contract.ID, creationDate.Unix())
	if err != nil {
		fmt.Println("createBills_1: ", err.Error())
		return
	}

	b := &bill.Bill{
		ContractId:         contract.ID,
		CreationDate:       creationDate.Unix(),
		ExpirationDate:     expirationDate,
		Paid:               false,
		Value:              value,
		DaysPastExpiration: 0,
		NeedsReconnection:  false,
	}

	created, err := b.WasAlreadyCreated()
	if err != nil {
		fmt.Println("createBills_2: ", err.Error())
		return
	}

	if created {
		return
	}

	fmt.Println("generating new bill...")

	_, err = b.Store()
	if err != nil {
		fmt.Println("createBills_3: ", err.Error())
		return
	}

	bills, err := bill.LoadAllPreviousTotals(contract.ID)
	if err != nil {
		fmt.Println("createBills_4: ", err.Error())
		return
	}

	body, err := reports.GetPDF(contract, bills[0], bills[1:])
	if err != nil {
		fmt.Println("createBills_5: ", err.Error())
		return
	}

	err = email.SendEmail(body)
	if err != nil {
		fmt.Println("createBills_6: ", err.Error())
		return
	}
}

func createReadingsAndBills() {
	contracts, err := contract.ListActiveContractsIDs()
	if err != nil {
		fmt.Println("createReadingsAndBills_1: ", err.Error())
		return
	}

	for _, contract := range contracts {
		createReadings(contract.ID)
		createBills(contract)
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

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        handler.NewRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("sending readings...")
				createReadingsAndBills()
			}
		}
	}()

	fmt.Println("Starting server")
	log.Fatal(server.ListenAndServe())
}
