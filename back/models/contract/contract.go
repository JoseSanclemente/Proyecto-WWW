package contract

import (
	"fmt"

	"Proyecto-WWW/back/shared/random"
	"Proyecto-WWW/back/storage"
)

type Contract struct {
	ID               string `json:"id"`
	ConsumerID       string `json:"consumer_id"`
	TransformerID    string `json:"transformer_id"`
	Address          string `json:"address"`
	NotificationType string `json:"notification_type"`
	Deleted          bool   `json:"deleted"`
}

func Load(id string) (*Contract, error) {
	rows, err := storage.DB.Query(
		"SELECT consumer, transformer, address, notification_type, deleted FROM contract WHERE id=?",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	if !rows.Next() {
		return nil, nil
	}

	contract := &Contract{
		ID: id,
	}
	err = rows.Scan(&contract.ConsumerID, &contract.TransformerID, &contract.Address, &contract.NotificationType, &contract.Deleted)
	if err != nil {
		return nil, err
	}

	return contract, nil
}

func List(consumerID string) ([]*Contract, error) {
	rows, err := storage.DB.Query(
		"SELECT id, transformer, address, notification_type, deleted FROM contract WHERE consumer = ?",
		consumerID,
	)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var contracts []*Contract
	for rows.Next() {
		c := &Contract{
			ConsumerID: consumerID,
		}

		err = rows.Scan(&c.ID, &c.TransformerID, &c.Address, &c.NotificationType, &c.Deleted)
		if err != nil {
			return nil, err
		}

		contracts = append(contracts, c)
	}

	return contracts, nil
}

func ListActiveContractsIDs() ([]*Contract, error) {
	rows, err := storage.DB.Query(
		"SELECT id, consumer FROM contract WHERE deleted=false",
	)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var contracts []*Contract
	for rows.Next() {
		c := &Contract{}

		err = rows.Scan(&c.ID, &c.ConsumerID)
		if err != nil {
			return nil, err
		}

		contracts = append(contracts, c)
	}

	return contracts, nil
}

func (c *Contract) Store() (string, error) {
	id := random.GenerateID("CON")

	result, err := storage.DB.Exec(
		"INSERT INTO contract(id, consumer, transformer, address, notification_type) VALUES(?, ?, ?, ?, ?)",
		id,
		c.ConsumerID,
		c.TransformerID,
		c.Address,
		c.NotificationType,
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

func (c *Contract) Update(notificationType, deleted string) error {
	if notificationType != "" {
		c.NotificationType = notificationType
	}

	if deleted == "true" {
		c.Deleted = true
	}

	if deleted == "false" {
		c.Deleted = false
	}

	result, err := storage.DB.Exec(
		"UPDATE contract SET notification_type=?, deleted=? WHERE id=?",
		c.NotificationType,
		c.Deleted,
		c.ID,
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
