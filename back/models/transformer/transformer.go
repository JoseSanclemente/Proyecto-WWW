package transformer

import (
	"fmt"

	"univalle/www/shared/random"
	"univalle/www/storage"
)

type Transformer struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	SubstationID string `json:"substation_id"`
	Longitude    string `json:"longitude"`
	Latitude     string `json:"latitude"`
	Deleted      bool   `json:"deleted"`
}

func Load(id string) (*Transformer, error) {
	rows, err := storage.DB.Query(
		"SELECT name, substation, longitude, latitude, deleted FROM transformer WHERE id=? AND deleted=FALSE",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	if !rows.Next() {
		return nil, nil
	}

	transformer := &Transformer{
		ID: id,
	}
	err = rows.Scan(&transformer.Name, &transformer.SubstationID, &transformer.Longitude, &transformer.Latitude, &transformer.Deleted)
	if err != nil {
		return nil, err
	}

	return transformer, nil
}

func List(substationID string) ([]*Transformer, error) {
	rows, err := storage.DB.Query(
		"SELECT id, name, longitude, latitude, deleted FROM transformer WHERE substation=?",
		substationID,
	)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var transformers []*Transformer
	for rows.Next() {
		transformer := &Transformer{
			SubstationID: substationID,
		}

		err = rows.Scan(&transformer.ID, &transformer.Name, &transformer.Longitude, &transformer.Latitude, &transformer.Deleted)
		if err != nil {
			return nil, err
		}

		transformers = append(transformers, transformer)
	}

	return transformers, nil
}

func (t *Transformer) Store() (string, error) {
	id := random.GenerateID("TRA")

	result, err := storage.DB.Exec(
		"INSERT INTO transformer(id, name, substation, longitude, latitude) VALUES(?, ?, ?, ?, ?)",
		id,
		t.Name,
		t.SubstationID,
		t.Longitude,
		t.Latitude,
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

func (t *Transformer) Update(name, longitude, latitude, deleted string) error {
	if name != "" {
		t.Name = name
	}

	if longitude != "" {
		t.Longitude = longitude
	}

	if latitude != "" {
		t.Latitude = latitude
	}

	if deleted == "true" {
		t.Deleted = true
	}

	if deleted == "false" {
		t.Deleted = false
	}

	result, err := storage.DB.Exec(
		"UPDATE transformer SET name=?, longitude=?, latitude=?, deleted=? WHERE id=?",
		t.Longitude,
		t.Latitude,
		t.Deleted,
		t.ID,
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
