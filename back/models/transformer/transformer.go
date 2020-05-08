package transformer

import (
	"fmt"

	"Proyecto-WWW/back/shared/random"
	"Proyecto-WWW/back/storage"
)

type Transformer struct {
	ID           string `json:"id"`
	SubstationID string `json:"substation_id"`
	Longitude    string `json:"longitude"`
	Latitude     string `json:"latitude"`
	Deleted      bool   `json:"deleted"`
}

func Load(id string) (*Transformer, error) {
	rows, err := storage.DB.Query(
		"SELECT substation, longitude, latitude, deleted FROM transformer WHERE id=? AND deleted=FALSE",
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
	err = rows.Scan(&transformer.SubstationID, &transformer.Longitude, &transformer.Latitude, &transformer.Deleted)
	if err != nil {
		return nil, err
	}

	return transformer, nil
}

func List(substationID string) ([]*Transformer, error) {
	rows, err := storage.DB.Query(
		"SELECT id, longitude, latitude, deleted FROM transformer WHERE substation=?",
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

		err = rows.Scan(&transformer.ID, &transformer.Longitude, &transformer.Latitude, &transformer.Deleted)
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
		"INSERT INTO transformer(id, substation, longitude, latitude) VALUES(?, ?, ?, ?)",
		id,
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

func (t *Transformer) Update(longitude, latitude, deleted string) error {
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
		"UPDATE transformer SET longitude=?, latitude=?, deleted=? WHERE id=?",
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
