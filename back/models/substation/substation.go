package substation

import (
	"fmt"

	"Proyecto-WWW/back/models/transformer"
	"Proyecto-WWW/back/shared/random"
	"Proyecto-WWW/back/storage"
)

type Substation struct {
	ID           string                     `json:"id"`
	Name         string                     `json:"name"`
	Longitude    string                     `json:"longitude"`
	Latitude     string                     `json:"latitude"`
	Deleted      bool                       `json:"deleted"`
	Transformers []*transformer.Transformer `json:"transformers"`
}

func Load(id string) (*Substation, error) {
	rows, err := storage.DB.Query(
		"SELECT name, longitude, latitude, deleted FROM substation WHERE id=? AND deleted=FALSE",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	if !rows.Next() {
		return nil, nil
	}

	substation := &Substation{
		ID: id,
	}
	err = rows.Scan(&substation.Name, &substation.Longitude, &substation.Latitude, &substation.Deleted)
	if err != nil {
		return nil, err
	}

	return substation, nil
}

func LoadAny(id string) (*Substation, error) {
	rows, err := storage.DB.Query(
		"SELECT name, longitude, latitude, deleted FROM substation WHERE id=?",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	if !rows.Next() {
		return nil, nil
	}

	substation := &Substation{
		ID: id,
	}
	err = rows.Scan(&substation.Name, &substation.Longitude, &substation.Latitude, &substation.Deleted)
	if err != nil {
		return nil, err
	}

	return substation, nil
}

func List() ([]*Substation, error) {
	rows, err := storage.DB.Query(
		"SELECT id, name, longitude, latitude, deleted FROM substation",
	)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var substations []*Substation
	for rows.Next() {
		substation := &Substation{}

		err = rows.Scan(&substation.ID, &substation.Name, &substation.Longitude, &substation.Latitude, &substation.Deleted)
		if err != nil {
			return nil, err
		}

		transformers, err := transformer.List(substation.ID)
		if err != nil {
			return nil, err
		}

		substation.Transformers = transformers
		substations = append(substations, substation)
	}

	return substations, nil
}

func (s *Substation) Store() (string, error) {
	id := random.GenerateID("SUB")

	result, err := storage.DB.Exec(
		"INSERT INTO substation(id, name, longitude, latitude) VALUES(?, ?, ?, ?)",
		id,
		s.Name,
		s.Longitude,
		s.Latitude,
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

func (s *Substation) Update(name, longitude, latitude, deleted string) error {
	if name != "" {
		s.Name = name
	}

	if longitude != "" {
		s.Longitude = longitude
	}

	if latitude != "" {
		s.Latitude = latitude
	}

	if deleted == "true" {
		s.Deleted = true
	}

	if deleted == "false" {
		s.Deleted = false
	}

	result, err := storage.DB.Exec(
		"UPDATE substation SET name=?, longitude=?, latitude=?, deleted=? WHERE id=?",
		s.Name,
		s.Longitude,
		s.Latitude,
		s.Deleted,
		s.ID,
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
