package bill

import (
	"Proyecto-WWW/back/shared/random"
	"Proyecto-WWW/back/storage"
	"fmt"
)

type Bill struct {
	ID             string `json:"id"`
	ContractId     string `json:"contract_id"`
	CreationDate   int64  `json:"creation_date"`
	ExpirationDate int64  `json:"expiration_date"`
	Paid           bool   `json:"paid"`
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
