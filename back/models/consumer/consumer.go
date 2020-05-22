package consumer

import (
	"fmt"

	"Proyecto-WWW/back/storage"
)

type Input struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Deleted  bool   `json:"deleted"`
}

type Consumer struct {
	ID       string `json:"id"`
	Password string `json:"-"`
	Email    string `json:"email"`
	Deleted  bool   `json:"deleted"`
}

func Login(id, password string) (bool, error) {
	rows, err := storage.DB.Query(
		"SELECT 1 FROM consumer WHERE id=? AND password=? AND deleted=FALSE",
		id,
		password,
	)
	if err != nil {
		return false, err
	}
	defer func() { _ = rows.Close() }()

	if !rows.Next() {
		return false, nil
	}

	return true, nil
}

func Load(id string) (*Consumer, error) {
	rows, err := storage.DB.Query(
		"SELECT password, email, deleted FROM consumer WHERE id=?",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	if !rows.Next() {
		return nil, nil
	}

	consumer := &Consumer{
		ID: id,
	}
	err = rows.Scan(&consumer.Password, &consumer.Email, &consumer.Deleted)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}

func List() ([]*Consumer, error) {
	rows, err := storage.DB.Query(
		"SELECT id, email, deleted FROM consumer",
	)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var consumers []*Consumer
	for rows.Next() {
		consumer := &Consumer{}

		err = rows.Scan(&consumer.ID, &consumer.Email, &consumer.Deleted)
		if err != nil {
			return nil, err
		}

		consumers = append(consumers, consumer)
	}

	return consumers, nil
}

func (c *Consumer) Store() error {
	result, err := storage.DB.Exec(
		"INSERT INTO consumer(id, password, email) VALUES(?, ?, ?)",
		c.ID,
		c.Password,
		c.Email,
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

func (c *Consumer) Update(password, email, deleted string) error {
	if password != "" {
		c.Password = password
	}

	if email != "" {
		c.Email = email
	}

	if deleted == "true" {
		c.Deleted = true
	}

	if deleted == "false" {
		c.Deleted = false
	}

	result, err := storage.DB.Exec(
		"UPDATE consumer SET password=?, email=?, deleted=? WHERE id=?",
		c.Password,
		c.Email,
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
