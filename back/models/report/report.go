package report

import "Proyecto-WWW/back/models/report/pdf"

func GetPDF() {
	_, _ = pdf.GetPDF(nil, nil)
}
