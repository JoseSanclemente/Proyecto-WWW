package pdf

import (
	"bytes"
	"encoding/base64"
	"strconv"
	"time"

	"univalle/www/models/bill"
	"univalle/www/models/contract"

	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

var (
	fontHelvetica = model.NewStandard14FontMustCompile(model.HelveticaName)

	white     = creator.ColorRGBFrom8bit(255, 255, 255)
	lightBlue = creator.ColorRGBFrom8bit(217, 240, 250)
	blue      = creator.ColorRGBFrom8bit(2, 136, 209)

	unitValue = 567

	monthNames = map[int]string{
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
)

func setBill(c *creator.Creator, contract *contract.Contract, bill *bill.Bill) error {
	logo, err := c.NewImageFromFile("models/report/pdf/assets/images/biplane.jpg")
	if err != nil {
		return err
	}

	invoice := c.NewInvoice()

	// Set invoice title.
	invoice.SetTitle("Recibo de energía")

	// Customize invoice title style.
	titleStyle := invoice.TitleStyle()
	titleStyle.Color = blue
	titleStyle.Font = fontHelvetica
	titleStyle.FontSize = 30

	invoice.SetTitleStyle(titleStyle)

	// Set invoice logo.
	invoice.SetLogo(logo)

	// Set invoice information.
	titleCell, _ := invoice.SetNumber(contract.ID)
	titleCell.Value = "Contrato"

	titleCell, _ = invoice.SetDate(time.Unix(bill.CreationDate, 0).In(time.UTC).Format("2006-01-02"))
	titleCell.Value = "Fecha de creación"

	titleCell, _ = invoice.SetDueDate(time.Unix(bill.ExpirationDate, 0).In(time.UTC).Format("2006-01-02"))
	titleCell.Value = "Fecha de vencimiento"

	// Customize invoice information styles.
	for _, info := range invoice.InfoLines() {
		descCell, contentCell := info[0], info[1]
		descCell.BackgroundColor = lightBlue
		contentCell.TextStyle.Font = fontHelvetica
	}

	// Set invoice addresses.
	invoice.SetSellerAddress(&creator.InvoiceAddress{
		Name:   "Cédula:    " + contract.ConsumerID,
		Street: "Dirección: " + contract.Address,
	})

	invoice.SetBuyerAddress(&creator.InvoiceAddress{})

	// Customize address styles.
	addressStyle := invoice.AddressStyle()
	addressStyle.Font = fontHelvetica
	addressStyle.FontSize = 9

	addressHeadingStyle := invoice.AddressHeadingStyle()
	addressHeadingStyle.Color = blue
	addressHeadingStyle.Font = fontHelvetica
	addressHeadingStyle.FontSize = 16

	invoice.SetAddressStyle(addressStyle)
	invoice.SetAddressHeadingStyle(addressHeadingStyle)

	tStyle := creator.InvoiceCellProps{
		TextStyle:       invoice.TitleStyle(),
		Alignment:       creator.CellHorizontalAlignmentLeft,
		BackgroundColor: white,
		BorderColor:     white,
		BorderWidth:     1,
		BorderSides:     []creator.CellBorderSide{creator.CellBorderSideAll},
	}
	// Customize column styles.
	invoice.SetColumns([]*creator.InvoiceCell{
		{
			InvoiceCellProps: tStyle,
			Value:            "Conceptos",
		},
		{
			InvoiceCellProps: tStyle,
			Value:            "Cantidad",
		},
		{
			InvoiceCellProps: tStyle,
			Value:            "Valor unitario",
		},
		{
			InvoiceCellProps: tStyle,
			Value:            "Valor total",
		},
	})

	for _, column := range invoice.Columns() {
		column.BackgroundColor = lightBlue
		column.BorderColor = lightBlue
		column.TextStyle.FontSize = 9
	}

	total := bill.Value * unitValue

	cells := invoice.AddLine(
		"Valor consumo energía",
		strconv.Itoa(bill.Value),
		strconv.Itoa(unitValue),
		"$"+strconv.Itoa(total),
	)

	for _, cell := range cells {
		cell.BorderColor = white
		cell.TextStyle.FontSize = 9
	}

	if bill.DaysPastExpiration > 0 {
		extra := int(float64(total) * (float64(bill.DaysPastExpiration) / 100))
		total += extra
		cells = invoice.AddLine(
			"Interes por mora",
			strconv.Itoa(bill.DaysPastExpiration)+"%",
			"",
			"$"+strconv.Itoa(extra),
		)

		for _, cell := range cells {
			cell.BorderColor = white
			cell.TextStyle.FontSize = 9
		}
	}

	if bill.NeedsReconnection {
		total += 34000
		cells = invoice.AddLine(
			"Reconexión",
			"1",
			"$34000",
			"$34000",
		)

		for _, cell := range cells {
			cell.BorderColor = white
			cell.TextStyle.FontSize = 9
		}
	}

	// Customize total line styles.
	titleCell, contentCell := invoice.Total()
	titleCell.BackgroundColor = lightBlue
	titleCell.BorderColor = lightBlue
	contentCell.BackgroundColor = lightBlue
	contentCell.BorderColor = lightBlue

	invoice.SetTotal("$" + strconv.Itoa(total))

	// Customize note styles.
	noteStyle := invoice.NoteStyle()
	noteStyle.Font = fontHelvetica
	noteStyle.FontSize = 12

	noteHeadingStyle := invoice.NoteHeadingStyle()
	noteHeadingStyle.Color = blue
	noteHeadingStyle.Font = fontHelvetica
	noteHeadingStyle.FontSize = 14

	invoice.SetNoteStyle(noteStyle)
	invoice.SetNoteHeadingStyle(noteHeadingStyle)

	if err = c.Draw(invoice); err != nil {
		return err
	}

	return nil
}

func setPreviousBills(c *creator.Creator, bills []*bill.Bill) error {
	in := c.NewInvoice()

	in.SetTitle("")

	// Customize invoice title style.
	inTitleStyle := in.TitleStyle()
	inTitleStyle.Color = blue
	inTitleStyle.Font = fontHelvetica
	inTitleStyle.FontSize = 1

	in.SetTitleStyle(inTitleStyle)

	in.BuyerAddress().Heading = "Consumos anteriores (kWh)"

	inAddressStyle := in.AddressStyle()
	inAddressStyle.Font = fontHelvetica
	inAddressStyle.FontSize = 1
	inAddressStyle.Color = blue
	in.SetAddressStyle(inAddressStyle)

	inAddress := in.AddressHeadingStyle()
	inAddress.Font = fontHelvetica
	inAddress.FontSize = 16
	in.SetAddressHeadingStyle(inAddress)

	ttStyle := creator.InvoiceCellProps{
		TextStyle:       in.TitleStyle(),
		Alignment:       creator.CellHorizontalAlignmentCenter,
		BackgroundColor: white,
		BorderColor:     white,
		BorderWidth:     1,
		BorderSides:     []creator.CellBorderSide{creator.CellBorderSideAll},
	}

	var months []*creator.InvoiceCell
	var values []string

	for _, b := range bills {
		d := time.Unix(b.CreationDate, 0).In(time.UTC)

		m := monthNames[int(d.AddDate(0, -1, 0).Month())]

		months = append(months, &creator.InvoiceCell{
			InvoiceCellProps: ttStyle,
			Value:            m,
		})
		values = append(values, strconv.Itoa(b.Value))
	}

	in.SetColumns(months)

	for _, column := range in.Columns() {
		column.BackgroundColor = lightBlue
		column.BorderColor = lightBlue
		column.TextStyle.FontSize = 9
	}

	cells := in.AddLine(values...)
	for _, cell := range cells {
		cell.BorderColor = white
		cell.TextStyle.FontSize = 9
	}

	if err := c.Draw(in); err != nil {
		return err
	}

	return nil
}

func setAd(c *creator.Creator) error {
	ad, err := c.NewImageFromFile("models/report/pdf/ads/ad.png")
	if err != nil {
		return err
	}

	ad.SetHorizontalAlignment(creator.HorizontalAlignmentCenter)
	if ad.Height() > 120 {
		ad.ScaleToHeight(120)
	}

	if ad.Width() > 500 {
		ad.ScaleToWidth(500)
	}

	if err = c.Draw(ad); err != nil {
		return err
	}

	return nil
}

func GetPDF(contract *contract.Contract, bill *bill.Bill, bills []*bill.Bill) ([]byte, error) {
	c := creator.New()
	c.NewPage()

	err := setBill(c, contract, bill)
	if err != nil {
		return nil, err
	}

	if len(bills) != 0 {
		err = setPreviousBills(c, bills)
		if err != nil {
			return nil, err
		}
	}

	err = setAd(c)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer([]byte{})

	err = c.Write(buf)
	if err != nil {
		return nil, err
	}

	pdf := base64.StdEncoding.EncodeToString(buf.Bytes())
	return []byte(pdf), nil
}
