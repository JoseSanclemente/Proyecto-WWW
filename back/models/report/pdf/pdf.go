package pdf

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"Proyecto-WWW/back/models/bill"
	"Proyecto-WWW/back/models/contract"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

var monthNames = map[int]string{
	1:  "Enero",
	2:  "Febrero",
	3:  "Marzo",
	4:  "Abril",
	5:  "Mayo",
	6:  "Junio",
	7:  "Julio",
	8:  "Agosto",
	9:  "Septiembre",
	10: "Octubre",
	11: "Noviembre",
	12: "Diciembre",
}

func GetPDF(contract *contract.Contract, bill *bill.Bill, bills []*bill.Bill) ([]byte, error) {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 15, 10)

	byteSlices, err := ioutil.ReadFile("../models/report/pdf/ads/ad.png")
	if err != nil {
		fmt.Println("Got error while opening file:", err)
		os.Exit(1)
	}

	ad := base64.StdEncoding.EncodeToString(byteSlices)

	byteSlices, err = ioutil.ReadFile("../models/report/pdf/assets/images/biplane.jpg")
	if err != nil {
		fmt.Println("Got error while opening file:", err)
		os.Exit(1)
	}

	base64 := base64.StdEncoding.EncodeToString(byteSlices)

	m.RegisterHeader(func() {
		m.Row(22, func() {
			m.Col(3, func() {
				err = m.Base64Image(base64, consts.Jpg, props.Rect{
					Center:  true,
					Percent: 70,
				})
			})

			m.ColSpace(3)

			m.SetBorder(true)
			m.Col(6, func() {
				m.Text(fmt.Sprintf("  Contrato:                              %s", bill.ContractId), props.Text{
					Size:   10,
					Style:  consts.Bold,
					Top:    1,
					Family: consts.Arial,
				})
				m.Text(fmt.Sprintf("  Cédula:                                 %s", contract.ConsumerID), props.Text{
					Size:   10,
					Style:  consts.Bold,
					Top:    5,
					Family: consts.Arial,
				})
				m.Text(fmt.Sprintf("  Dirección:                             %s", contract.Address), props.Text{
					Size:   10,
					Style:  consts.Bold,
					Top:    9,
					Family: consts.Arial,
				})
				m.Text(fmt.Sprintf("  Fecha de creacion:             %s", time.Unix(bill.CreationDate, 0).In(time.UTC).Format("2006-01-02")), props.Text{
					Size:   10,
					Style:  consts.Bold,
					Top:    13,
					Family: consts.Arial,
				})
				m.Text(fmt.Sprintf("  Fecha de vencimiento:        %s", time.Unix(bill.ExpirationDate, 0).In(time.UTC).Format("2006-01-02")), props.Text{
					Size:   10,
					Style:  consts.Bold,
					Top:    17,
					Family: consts.Arial,
				})
			})
			m.SetBorder(false)
		})

		m.Line(5.0)
	})
	if err != nil {
		return nil, err
	}

	m.Row(1, func() {})

	m.Row(5, func() {
		m.Col(0, func() {
			m.Text(fmt.Sprintf("Consumos anteriores (kWh)"), props.Text{
				Top:    0,
				Family: consts.Arial,
				Style:  consts.Bold,
			})
		})
	})

	var months []string
	values := [][]string{{}}

	for _, b := range bills {
		d := time.Unix(b.CreationDate, 0).In(time.UTC)

		m := monthNames[int(d.AddDate(0, -1, 0).Month())]

		months = append(months, m)
		values[0] = append(values[0], strconv.Itoa(b.Value))
	}

	m.TableList(months, values, props.TableList{
		HeaderContentSpace: 1,
		Align:              consts.Center,
	})

	m.Line(10.0)

	m.Row(15, func() {
		m.ColSpace(2)
		m.Col(3, func() {
			m.Text(fmt.Sprintf("  Consumo:         %d", bill.Value), props.Text{
				Size:   10,
				Style:  consts.Bold,
				Top:    1,
				Family: consts.Arial,
			})
			m.Text("  Valor unitario:  123456789", props.Text{
				Size:   10,
				Style:  consts.Bold,
				Top:    5,
				Family: consts.Arial,
			})
			m.Text(fmt.Sprintf("  Tota a pagar:    %d", bill.Value), props.Text{
				Size:   10,
				Style:  consts.Bold,
				Top:    9,
				Family: consts.Arial,
			})
		})

		m.ColSpace(1)

		m.Col(6, func() {
			m.Text("Pendiente por pagar:", props.Text{
				Size:   10,
				Style:  consts.Bold,
				Top:    1,
				Family: consts.Arial,
				Align:  consts.Center,
			})
			m.Text(strconv.Itoa(bill.Value), props.Text{
				Size:   10,
				Style:  consts.Bold,
				Top:    5,
				Family: consts.Arial,
				Align:  consts.Center,
			})
		})
	})

	m.Row(10, func() {})

	m.Row(30, func() {
		m.Col(0, func() {
			err = m.Base64Image(ad, consts.Png, props.Rect{
				Center:  true,
				Percent: 100,
			})
		})
	})
	if err != nil {
		return nil, err
	}

	buf, err := m.Output()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
