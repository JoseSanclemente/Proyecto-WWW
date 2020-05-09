package bill

import (
	"Proyecto-WWW/back/storage"
	"fmt"
	"time"
)

type Bill struct {
	ID             string
	ContractId     string
	CreationDate   time.Time
	ExpirationDate time.Time
	Payed          bool
}

func Load(id string) (*Bill, error) {
	rows, err := storage.DB.Query(
		"SELECT contract, creation_date, expiration_date, payed FROM bill WHERE id=?",
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
	err = rows.Scan(&bill.ContractId, &bill.CreationDate, &bill.ExpirationDate, &bill.Payed)
	if err != nil {
		return nil, err
	}

	return bill, nil
}

func (b *Bill) RegisterPayment() error {
	result, err := storage.DB.Exec(
		"UPDATE bill SET payed=true WHERE id=?",
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

func (b *Bill) Store() error {
	result, err := storage.DB.Exec(
		"INSERT INTO bill(id, contract, creation_date, expiration_date, payed) VALUES(?, ?, ?, ?, false)",
		b.ID,
		b.ContractId,
		b.CreationDate,
		b.ExpirationDate,
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
