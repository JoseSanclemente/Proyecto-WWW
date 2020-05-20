package consumer

import (
	"errors"
	"testing"

	"univalle/www/storage"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	c := require.New(t)

	db, mock, err := sqlmock.New()
	c.Nil(err)
	defer db.Close()

	storage.DB = db

	mock.ExpectQuery("SELECT 1 FROM consumer").
		WillReturnRows(sqlmock.NewRows([]string{"id", "password"}).AddRow("id", "password"))

	ok, err := Login("id", "password")
	c.Nil(err)
	c.True(ok)
	c.Nil(mock.ExpectationsWereMet())

	mock.ExpectQuery("SELECT 1 FROM consumer").
		WillReturnRows(sqlmock.NewRows([]string{"id", "password"}))

	ok, err = Login("id", "password")
	c.Nil(err)
	c.False(ok)
	c.Nil(mock.ExpectationsWereMet())
}

func TestLoginError(t *testing.T) {
	c := require.New(t)

	db, mock, err := sqlmock.New()
	c.Nil(err)
	defer db.Close()

	storage.DB = db

	e := errors.New("query error")
	mock.ExpectQuery("SELECT 1 FROM consumer").WillReturnError(e)

	_, err = Login("id", "password")
	c.Equal(e, err)
}

func TestLoad(t *testing.T) {
	c := require.New(t)

	db, mock, err := sqlmock.New()
	c.Nil(err)
	defer db.Close()

	storage.DB = db

	mock.ExpectQuery("SELECT password, email, deleted FROM consumer").
		WillReturnRows(sqlmock.NewRows([]string{"password", "email", "deleted"}).
			AddRow("password", "email", false))

	consumer, err := Load("id")
	c.Nil(err)
	c.Nil(mock.ExpectationsWereMet())
	c.Equal("email", consumer.Email)

	mock.ExpectQuery("SELECT password, email, deleted FROM consumer").
		WillReturnRows(sqlmock.NewRows([]string{"password", "email", "deleted"}))

	consumer, err = Load("id")
	c.Nil(err)
	c.Nil(consumer)
	c.Nil(mock.ExpectationsWereMet())
}

func TestLoadError(t *testing.T) {
	c := require.New(t)

	db, mock, err := sqlmock.New()
	c.Nil(err)
	defer db.Close()

	storage.DB = db

	e := errors.New("query error")
	mock.ExpectQuery("SELECT password, email, deleted FROM consumer").WillReturnError(e)

	_, err = Load("id")
	c.Equal(e, err)
}

func TestList(t *testing.T) {
	c := require.New(t)

	db, mock, err := sqlmock.New()
	c.Nil(err)
	defer db.Close()

	storage.DB = db

	mock.ExpectQuery("SELECT id, email, deleted FROM consumer").
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "deleted"}).
			AddRow("id", "email", false))

	consumers, err := List()
	c.Nil(err)
	c.Nil(mock.ExpectationsWereMet())
	c.Equal("email", consumers[0].Email)
}

func TestStore(t *testing.T) {
	c := require.New(t)

	db, mock, err := sqlmock.New()
	c.Nil(err)
	defer db.Close()

	storage.DB = db
	consumer := &Consumer{
		ID:       "id",
		Password: "password",
		Email:    "email",
	}

	mock.ExpectExec("INSERT INTO consumer").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = consumer.Store()
	c.Nil(err)
	c.Nil(mock.ExpectationsWereMet())
}

func TestUpdate(t *testing.T) {
	c := require.New(t)

	db, mock, err := sqlmock.New()
	c.Nil(err)
	defer db.Close()

	storage.DB = db
	consumer := &Consumer{
		ID:       "id",
		Password: "password",
		Email:    "email",
	}

	mock.ExpectExec("UPDATE consumer").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = consumer.Update("password", "email", "false")
	c.Nil(err)
	c.Nil(mock.ExpectationsWereMet())
}
