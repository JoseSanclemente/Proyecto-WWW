package reading

import (
	"fmt"
	"time"

	"Proyecto-WWW/back/shared/random"
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
	to := int(date.Add(time.Hour * 24 * 31).Unix())

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

func (r *Reading) Store() error {
	id := random.GenerateID("REA")

	result, err := storage.DB.Exec(
		"INSERT INTO reading(id, contract, value, date) VALUES(?, ?, ?, ?)",
		id,
		r.ContractID,
		r.Value,
		r.Date,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return fmt.Errorf("expected to affect 1 row, affected %d", rows)
	}

	return nil
}
