package report

import (
	"Proyecto-WWW/back/models/bill"
	"Proyecto-WWW/back/models/contract"
	"Proyecto-WWW/back/models/report/pdf"
)

func GetPDF(contract *contract.Contract, bill *bill.Bill, bills []*bill.Bill) ([]byte, error) {
	return pdf.GetPDF(contract, bill, bills)
}

func GetMostConsumedMonth() {

}

func GetMostConsumedMonths() {

}

func GetMostConsumedMonthByYear() {

}

func GetLastMonthsConsume() {

}

func GetTopConsumers() {

}

func GetTopConsumersExpiredContracts() {

}
