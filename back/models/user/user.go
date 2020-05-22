package user

import (
	"fmt"

	"Proyecto-WWW/back/storage"
)

type Input struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Type     string `json:"type"`
	Deleted  bool   `json:"deleted"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"-"`
	Type     string `json:"type"`
	Deleted  bool   `json:"deleted"`
}

func Login(email, password string) (string, error) {
	rows, err := storage.DB.Query(
		"SELECT type FROM employee WHERE email=? AND password=? AND deleted=FALSE",
		email,
		password,
	)
	if err != nil {
		return "", err
	}
	defer func() { _ = rows.Close() }()

	if !rows.Next() {
		return "", nil
	}

	var userType string
	err = rows.Scan(&userType)
	if err != nil {
		return "", err
	}

	return userType, nil
}

func Load(email string) (*User, error) {
	rows, err := storage.DB.Query(
		"SELECT password, type, deleted FROM employee WHERE email=?",
		email,
	)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	if !rows.Next() {
		return nil, nil
	}

	user := &User{
		Email: email,
	}
	err = rows.Scan(&user.Password, &user.Type, &user.Deleted)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func List() ([]*User, error) {
	rows, err := storage.DB.Query(
		"SELECT email, password, type, deleted FROM employee",
	)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var users []*User
	for rows.Next() {
		user := &User{}

		err = rows.Scan(&user.Email, &user.Password, &user.Type, &user.Deleted)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *User) Store() error {
	result, err := storage.DB.Exec(
		"INSERT INTO employee(email, password, type) VALUES(?, ?, ?)",
		u.Email,
		u.Password,
		u.Type,
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

func (u *User) Update(password, userType, deleted string) error {
	if password != "" {
		u.Password = password
	}

	if userType != "" {
		u.Type = userType
	}

	if deleted == "true" {
		u.Deleted = true
	}

	if deleted == "false" {
		u.Deleted = false
	}

	result, err := storage.DB.Exec(
		"UPDATE employee SET password=?, type=?, deleted=? WHERE email=?",
		u.Password,
		u.Type,
		u.Deleted,
		u.Email,
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
