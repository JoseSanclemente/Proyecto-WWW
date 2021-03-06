package report

import (
	"time"

	"Proyecto-WWW/back/models/bill"
	"Proyecto-WWW/back/models/consumer"
	"Proyecto-WWW/back/models/contract"
	"Proyecto-WWW/back/models/report/email"
	"Proyecto-WWW/back/models/report/pdf"
	"Proyecto-WWW/back/storage"
)

type MonthConsumption struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Value int `json:"value"`
}

type YearConsumption struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

type TopConsumer struct {
	ID    string `json:"id"`
	Value int    `json:"value"`
}

func GetPDF(contract *contract.Contract, bill *bill.Bill, bills []*bill.Bill) ([]byte, error) {
	return pdf.GetPDF(contract, bill, bills)
}

func GetMonthlyConsumption(startYear, startMonth, endYear, endMonth int) ([]*MonthConsumption, error) {
	start := time.Date(startYear, time.Month(startMonth), 0, 0, 0, 0, 0, time.UTC)
	end := time.Date(endYear, time.Month(endMonth+1), 0, 0, 0, 0, 0, time.UTC).Add(time.Second * -1).In(time.UTC)

	rows, err := storage.DB.Query(
		`SELECT value, date FROM reading WHERE date >= ? AND date < ? ORDER BY date ASC `,
		start.Unix(),
		end.Unix(),
	)

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	months := []*MonthConsumption{}

	var value int
	var date int

	for rows.Next() {
		err = rows.Scan(&value, &date)
		if err != nil {
			return nil, err
		}

		t := time.Unix(int64(date), 0).In(time.UTC)

		if t.Year() > startYear {
			startYear = t.Year()
			startMonth = 1

			months = append(months, &MonthConsumption{
				Year:  startYear,
				Month: startMonth,
			})
		} else if int(t.Month()) > startMonth {
			startMonth = int(t.Month())

			months = append(months, &MonthConsumption{
				Year:  startYear,
				Month: startMonth,
			})
		}

		if len(months) == 0 {
			months = append(months, &MonthConsumption{
				Year:  startYear,
				Month: startMonth,
				Value: value,
			})
		} else {
			months[len(months)-1].Value += value
		}
	}

	return months, nil
}

func GetYearlyConsumption(startYear, endYear int) ([]*YearConsumption, error) {
	start := time.Date(startYear, 0, 0, 0, 0, 0, 0, time.UTC)
	end := time.Date(endYear+1, 0, 0, 0, 0, 0, 0, time.UTC).Add(time.Second * -1).In(time.UTC)

	rows, err := storage.DB.Query(
		`SELECT value, date FROM reading WHERE date >= ? AND date < ? ORDER BY date ASC`,
		start.Unix(),
		end.Unix(),
	)

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	years := []*YearConsumption{{
		Year: startYear,
	}}

	var value int
	var date int

	for rows.Next() {
		err = rows.Scan(&value, &date)
		if err != nil {
			return nil, err
		}

		t := time.Unix(int64(date), 0).In(time.UTC)

		if t.Year() > startYear {
			startYear = t.Year()

			years = append(years, &YearConsumption{
				Year: startYear,
			})
		}

		if len(years) == 0 {
			years = append(years, &YearConsumption{
				Year:  startYear,
				Value: value,
			})
		} else {
			years[len(years)-1].Value += value
		}
	}

	return years, nil
}

func GetLastMonthsConsume() {

}

func GetTopConsumers() ([]*TopConsumer, error) {
	contracts := map[string]int{}

	rows, err := storage.DB.Query(
		`SELECT contract, SUM(value) FROM reading GROUP BY contract`,
	)
	if err != nil {
		return nil, err
	}

	var contract string
	var value int
	for rows.Next() {
		err = rows.Scan(&contract, &value)
		if err != nil {
			return nil, err
		}

		contracts[contract] = value
	}

	_ = rows.Close()
	consumers := map[string]int{}

	rows, err = storage.DB.Query(
		`SELECT id, consumer FROM contract`,
	)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var consumer string
	for rows.Next() {
		err = rows.Scan(&contract, &consumer)
		if err != nil {
			return nil, err
		}

		consumers[consumer] += contracts[contract]
	}

	top := make([]*TopConsumer, 0, len(consumers))
	for consumer, value = range consumers {
		if len(top) >= 10 {
			break
		}

		set := false
		for i := range top {
			if value > top[i].Value {
				top = append(top[:i],
					append([]*TopConsumer{{ID: consumer, Value: value}},
						top[i:]...)...)
				set = true
				break
			}
		}

		if !set {
			top = append(top, &TopConsumer{ID: consumer, Value: value})
		}
	}

	if len(top) > 10 {
		top = top[:10]
	}

	return top, nil
}

func GetTopConsumersExpiredContracts() {

}

func SendEmail(consumerID string, file []byte) error {
	con, err := consumer.Load(consumerID)
	if err != nil {
		return err
	}

	if con.Email == "" {
		return nil
	}

	return email.SendEmail(con.Email, file)
}
