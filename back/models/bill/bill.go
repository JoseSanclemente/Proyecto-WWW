package bill

import (
	"fmt"
	"time"

	"Proyecto-WWW/back/models/reading"
	"Proyecto-WWW/back/shared/random"
	"Proyecto-WWW/back/storage"
)

type Bill struct {
	ID             string `json:"id"`
	ContractId     string `json:"contract_id"`
	CreationDate   int64  `json:"creation_date"`
	ExpirationDate int64  `json:"expiration_date"`
	Paid           bool   `json:"paid"`
	Value          int    `json:"value"`
}

func Load(id string) (*Bill, error) {
	rows, err := storage.DB.Query(
		"SELECT contract, creation_date, expiration_date, paid FROM bill WHERE id=?",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	if !rows.Next() {
		return nil, nil
	}

	bill := &Bill{
		ID: id,
	}
	err = rows.Scan(&bill.ContractId, &bill.CreationDate, &bill.ExpirationDate, &bill.Paid)
	if err != nil {
		return nil, err
	}

	return bill, nil
}

func LoadUnpaid(contractID string) (*Bill, error) {
	rows, err := storage.DB.Query(
		"SELECT id, creation_date, expiration_date FROM bill WHERE contract = ? AND paid = false ORDER BY creation_date ASC",
		contractID,
	)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var bills []*Bill
	for rows.Next() {
		bill := &Bill{
			ContractId: contractID,
		}

		err = rows.Scan(&bill.ID, &bill.CreationDate, &bill.ExpirationDate)
		if err != nil {
			return nil, err
		}

		bills = append(bills, bill)
	}

	if len(bills) == 0 {
		return nil, nil
	}

	value, err := reading.LoadTotal(bills[0].ContractId, bills[0].CreationDate)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	expiration := time.Unix(bills[0].ExpirationDate, 0)

	diff := now.Sub(expiration)

	if diff > (time.Hour * 24 * 30) {
		value += int(float64(value) * 0.3)
	} else {
		value += value * int(diff/(time.Hour*24))
	}

	if len(bills) > 1 && bills[1].ExpirationDate > now.Unix() {
		value += 34000
	}

	bills[0].Value = value

	return bills[0], nil
}

func (b *Bill) RegisterPayment() error {
	result, err := storage.DB.Exec(
		"UPDATE bill SET paid=true WHERE id=?",
		b.ID,
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

func (b *Bill) Store() (string, error) {
	id := random.GenerateID("BIL")

	result, err := storage.DB.Exec(
		"INSERT INTO bill(id, contract, creation_date, expiration_date, paid) VALUES(?, ?, ?, ?, false)",
		id,
		b.ContractId,
		b.CreationDate,
		b.ExpirationDate,
	)
	if err != nil {
		return "", err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return "", err
	}

	if rows != 1 {
		return "", fmt.Errorf("expected to affect 1 row, affected %d", rows)
	}

	return id, nil
}
