package bill

import (
	"errors"
	"fmt"
	"time"

	"univalle/www/models/reading"
	"univalle/www/shared/random"
	"univalle/www/storage"
)

type Bill struct {
	ID                 string `json:"id"`
	ContractId         string `json:"contract_id"`
	CreationDate       int64  `json:"creation_date"`
	ExpirationDate     int64  `json:"expiration_date"`
	Paid               bool   `json:"paid"`
	Value              int
	DaysPastExpiration int
	NeedsReconnection  bool
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

	diff := int(now.Sub(expiration) / time.Hour * 24)
	if diff > 30 {
		diff = 30
	}
	bills[0].DaysPastExpiration = diff

	if len(bills) > 1 && bills[1].ExpirationDate > now.Unix() {
		bills[0].NeedsReconnection = true
	}

	bills[0].Value = value

	return bills[0], nil
}

func LoadPreviousTotals(contractID string) ([]*Bill, error) {
	rows, err := storage.DB.Query(
		"SELECT id, creation_date, expiration_date FROM bill WHERE contract = ? AND paid = true ORDER BY creation_date DESC  LIMIT 5",
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

		value, err := reading.LoadTotal(bill.ContractId, bill.CreationDate)
		if err != nil {
			return nil, err
		}

		bill.Value = value
		bills = append(bills, bill)
	}

	return bills, nil
}

func LoadAllPreviousTotals(contractID string) ([]*Bill, error) {
	rows, err := storage.DB.Query(
		"SELECT id, creation_date, expiration_date, paid FROM bill WHERE contract = ? ORDER BY creation_date DESC LIMIT 6",
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

		err = rows.Scan(&bill.ID, &bill.CreationDate, &bill.ExpirationDate, &bill.Paid)
		if err != nil {
			return nil, err
		}

		value, err := reading.LoadTotal(bill.ContractId, bill.CreationDate)
		if err != nil {
			return nil, err
		}

		bill.Value = value
		bills = append(bills, bill)
	}

	if len(bills) == 0 {
		return nil, errors.New("no bills found")
	}

	return bills, nil
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

func IsContractActive(contractID string) (bool, error) {
	rows, err := storage.DB.Query(
		"SELECT expiration_date FROM bill WHERE contract = ? AND paid = false ORDER BY creation_date ASC",
		contractID,
	)
	if err != nil {
		return true, err
	}
	defer func() { _ = rows.Close() }()

	var bills []*Bill
	for rows.Next() {
		bill := &Bill{
			ContractId: contractID,
		}

		err = rows.Scan(&bill.ExpirationDate)
		if err != nil {
			return true, err
		}

		bills = append(bills, bill)
	}

	if len(bills) < 2 {
		return true, nil
	}

	now := time.Now()
	expiration := time.Unix(bills[1].ExpirationDate, 0)

	diff := now.Sub(expiration)

	return diff < 0, nil
}

func (b *Bill) WasAlreadyCreated() (bool, error) {
	year, month, _ := time.Now().UTC().Date()
	date := time.Date(year, month, 0, 0, 0, 0, 0, time.UTC).Unix()

	rows, err := storage.DB.Query(
		"SELECT COUNT(*) FROM bill WHERE contract = ? AND creation_date >= ?",
		b.ContractId,
		date,
	)
	if err != nil {
		return false, err
	}
	defer func() { _ = rows.Close() }()

	if !rows.Next() {
		return false, nil
	}

	var count int
	err = rows.Scan(&count)
	if err != nil {
		return false, err
	}

	return count != 0, nil
}
