package handler

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"univalle/www/models/bill"
	"univalle/www/models/contract"
	reports "univalle/www/models/report"
	"univalle/www/shared/gateway"
	"univalle/www/shared/random"
)

func registerOperatorPayment(response http.ResponseWriter, request *http.Request) {
	params := map[string]string{}

	err := json.NewDecoder(request.Body).Decode(&params)
	if err != nil {
		fmt.Println("registerOperatorPayment_1: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	b, err := bill.LoadUnpaid(params["contract_id"])
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
	contractID := request.URL.Query().Get("contract_id")
	if contractID == "" {
		response.Header().Set("Access-Control-Allow-Origin", "*")
		gateway.WriteJSON(response, 400, map[string]string{"message": "missing contract id"})
		return
	}

	contract := &contract.Contract{
		ID:               "CON1234567890123456",
		ConsumerID:       "1234531",
		TransformerID:    "TRA1234567890123456",
		Address:          "Cra 84a #14-115",
		NotificationType: "Email",
		Deleted:          false,
	}

	// contract, err := contract.Load(params["contract_id"])
	// if err != nil {
	// 	fmt.Println("getPDF_2: ", err.Error())
	// 	gateway.WriteInternalServerError(response)
	// 	return
	// }

	bills := []*bill.Bill{
		{
			CreationDate: 1583020800,
			Value:        145,
		},
		{
			CreationDate: 1580515200,
			Value:        170,
		},
		{
			CreationDate: 1577836800,
			Value:        100,
		},
		{
			CreationDate: 1575158400,
			Value:        5,
		},
		{
			CreationDate: 1572566400,
			Value:        200,
		},
	}

	// bills, err := bill.LoadPreviousTotals(params["contract_id"])
	// if err != nil {
	// 	fmt.Println("getPDF_3: ", err.Error())
	// 	gateway.WriteInternalServerError(response)
	// 	return
	// }

	bill := &bill.Bill{
		ID:                 "BIL1234567890123456",
		ContractId:         "CON1234567890123456",
		CreationDate:       1585699200,
		ExpirationDate:     1588204800,
		Paid:               false,
		Value:              150,
		DaysPastExpiration: 30,
		NeedsReconnection:  true,
	}

	// bill, err := bill.LoadUnpaid(params["contract_id"])
	// if err != nil {
	// 	fmt.Println("getPDF_4: ", err.Error())
	// 	gateway.WriteInternalServerError(response)
	// 	return
	// }

	b, err := reports.GetPDF(contract, bill, bills)
	if err != nil {
		fmt.Println("getPDF_5: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}

	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Content-Disposition", "attachment; filename=recibo.pdf")
	response.Header().Set("Content-Type", "application/pdf")
	// response.Header().Set("Content-Length ", strconv.Itoa(len(b)))
	response.WriteHeader(200)

	_, err = response.Write(b)
	if err != nil {
		fmt.Println("getPDF_6: ", err.Error())
		gateway.WriteInternalServerError(response)
		return
	}
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
