package handler

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"Proyecto-WWW/back/models/bill"
	reports "Proyecto-WWW/back/models/report"
	"Proyecto-WWW/back/shared/gateway"
	"Proyecto-WWW/back/shared/random"
)

func registerOperatorPayment(response http.ResponseWriter, request *http.Request) {
	params := map[string]string{}

	err := json.NewDecoder(request.Body).Decode(&params)
	if err != nil {
		fmt.Println("registerOperatorPayment_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	b, err := bill.Load(params["id"])
	if err != nil {
		fmt.Println("registerOperatorPayment_2: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	if b == nil {
		gateway.WriteJSON(response, 400, map[string]string{"error": "bill not found"})
		return
	}

	err = b.RegisterPayment()
	if err != nil {
		fmt.Println("registerOperatorPayment_3: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	gateway.WriteJSON(response, 200, map[string]string{"message": "ok"})
}

func createBill(response http.ResponseWriter, request *http.Request) {
	newBill := &bill.Bill{}

	err := json.NewDecoder(request.Body).Decode(newBill)
	if err != nil {
		fmt.Println("createBill_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	id, err := newBill.Store()
	if err != nil {
		fmt.Println("createBill_2: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	gateway.WriteJSON(response, 200, map[string]string{"id": id})
}

func readBankFile(request *http.Request) ([][]string, error) {
	err := request.ParseMultipartForm(256)
	if err != nil {
		fmt.Println("readBankFile_1: ", err.Error())
		return nil, err
	}

	file, _, err := request.FormFile("file")
	if err != nil {
		fmt.Println("readBankFile_2: ", err.Error())
		return nil, err
	}
	defer file.Close()

	filename := random.GenerateID("FIL")

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("readBankFile_3: ", err.Error())
		return nil, err
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		fmt.Println("readBankFile_4: ", err.Error())
		return nil, err
	}

	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("readBankFile_5: ", err.Error())
		return nil, err
	}

	r := csv.NewReader(readFile)

	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("readBankFile_5: ", err.Error())
		return nil, err
	}

	return records, nil
}

func getPDF(response http.ResponseWriter, request *http.Request) {
	params := map[string]string{}

	err := json.NewDecoder(request.Body).Decode(&params)
	if err != nil {
		fmt.Println("getPDF_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	bills, err := bill.LoadPreviousTotals(params["contract_id"])
	if err != nil {
		fmt.Println("getPDF_2: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	bill, err := bill.LoadUnpaid(params["contract_id"])
	if err != nil {
		fmt.Println("getPDF_3: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	_ = bill
	_ = bills
	reports.GetPDF()

	gateway.WriteJSON(response, 200, map[string]string{"message": "ok"})
}

func registerBankPayments(response http.ResponseWriter, request *http.Request) {
	rows, err := readBankFile(request)
	if err != nil {
		fmt.Println("registerBankPayments_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	for _, row := range rows {
		b, err := bill.Load(row[0])
		if err != nil {
			fmt.Println("registerBankPayments_2: ", err.Error())
			gateway.WriteInternalServerError(response)
			return
		}

		if b == nil {
			gateway.WriteJSON(response, 400, map[string]string{"error": "bill not found"})
			return
		}

		err = b.RegisterPayment()
		if err != nil {
			fmt.Println("registerBankPayments_3: ", err.Error())
			gateway.WriteInternalServerError(response)
			return
		}
	}

	gateway.WriteJSON(response, 200, map[string]string{"message": "ok"})
}
