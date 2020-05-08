package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {
	begin := time.Now()
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 15, 10)
	//m.SetBorder(true)

	byteSlices, err := ioutil.ReadFile("ads/ad.png")
	if err != nil {
		fmt.Println("Got error while opening file:", err)
		os.Exit(1)
	}

	ad := base64.StdEncoding.EncodeToString(byteSlices)

	byteSlices, err = ioutil.ReadFile("assets/images/biplane.jpg")
	if err != nil {
		fmt.Println("Got error while opening file:", err)
		os.Exit(1)
	}

	base64 := base64.StdEncoding.EncodeToString(byteSlices)

	m.RegisterHeader(func() {
		m.Row(20, func() {
			m.Col(3, func() {
				_ = m.Base64Image(base64, consts.Jpg, props.Rect{
					Center:  true,
					Percent: 70,
				})
			})

			m.ColSpace(3)

			m.SetBorder(true)
			m.Col(6, func() {
				m.Text("  Contrato:                              12345678901234567890", props.Text{
					Size:   10,
					Style:  consts.Bold,
					Top:    1,
					Family: consts.Arial,
				})
				m.Text("  Cédula:                                 123456", props.Text{
					Size:   10,
					Style:  consts.Bold,
					Top:    5,
					Family: consts.Arial,
				})
				m.Text("  Dirección:                             123456", props.Text{
					Size:   10,
					Style:  consts.Bold,
					Top:    9,
					Family: consts.Arial,
				})
				m.Text("  Fecha de vencimiento:        123456", props.Text{
					Size:   10,
					Style:  consts.Bold,
					Top:    13,
					Family: consts.Arial,
				})
			})
			m.SetBorder(false)
		})

		m.Line(5.0)
	})

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

	meses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo"}
	consumos := [][]string{{"123456", "123456", "123456", "123456", "123456"}}

	m.TableList(meses, consumos, props.TableList{
		HeaderContentSpace: 1,
		Align:              consts.Center,
	})

	m.Line(10.0)

	m.Row(15, func() {
		m.ColSpace(2)
		m.Col(3, func() {
			m.Text("  Consumo:         123456789", props.Text{
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
			m.Text("  Tota a pagar:    123456789", props.Text{
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
			m.Text("123456", props.Text{
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
			_ = m.Base64Image(ad, consts.Png, props.Rect{
				Center:  true,
				Percent: 100,
			})
		})
	})

	err = m.OutputFileAndClose("sample1.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}
