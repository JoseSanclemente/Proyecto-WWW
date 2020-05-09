package reading

import (
	"time"

	"Proyecto-WWW/back/storage"
)

type Reading struct {
	ID         string
	ContractID string
	Value      int
	Date       int
}

func LoadTotal(contractID string, t int64) (int, error) {
	creation := time.Unix(t, 0)
	year, month, _ := creation.Date()
	date := time.Date(year, month, 0, 0, 0, 0, 0, creation.Location())

	from := int(date.AddDate(0, -1, 0).Unix())
	to := int(date.Add(time.Second * -1).Unix())

	rows, err := storage.DB.Query(
		"SELECT SUM(value) FROM reading WHERE contract=? AND date >= ? AND date < ?",
		contractID,
		from,
		to,
	)
	if err != nil {
		return 0, err
	}
	defer func() { _ = rows.Close() }()

	if !rows.Next() {
		return 0, nil
	}

	value := 0
	err = rows.Scan(&value)
	if err != nil {
		return 0, err
	}

	return value, nil
}
