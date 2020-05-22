package storage

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbFIle = "foo.db"
	dsn    = dbFIle + "?_foreign_keys=1"
)

var DB *sql.DB

func Connect() error {
	var err error
	DB, err = sql.Open("sqlite3", dsn)
	if err != nil {
		return err
	}

	return nil
}

func ResetAndConnect() error {
	err := os.Remove(dbFIle)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	err = Connect()
	if err != nil {
		return err
	}

	statement := `
	CREATE TABLE employee 
	(
		email VARCHAR(20) PRIMARY KEY NOT NULL,
		password VARCHAR(20) NOT NULL,
		type VARCHAR(20) NOT NULL,    
		deleted BOOLEAN DEFAULT 0
	);

	CREATE TABLE substation 
	(
		id VARCHAR(20) PRIMARY KEY NOT NULL,
		name VARCHAR(20) NOT NULL,
		longitude VARCHAR(20) NOT NULL,
		latitude VARCHAR(20) NOT NULL,
		deleted BOOLEAN DEFAULT 0
	);

	CREATE TABLE transformer 
	(
		id VARCHAR(20) PRIMARY KEY NOT NULL,
		name VARCHAR(20) NOT NULL,
		substation VARCHAR(20) NOT NULL,
		longitude VARCHAR(20) NOT NULL,
		latitude VARCHAR(20) NOT NULL,
		deleted BOOLEAN DEFAULT 0,
		
		CONSTRAINT fk_transformer_substation FOREIGN KEY (substation) REFERENCES substation(id)
	);

	CREATE TABLE consumer
	(
		id VARCHAR(20) PRIMARY KEY NOT NULL,
		password VARCHAR(20) NOT NULL,
		email VARCHAR(20) NOT NULL,
		deleted BOOLEAN DEFAULT 0
	);

	CREATE TABLE contract
	(
		id VARCHAR(20) PRIMARY KEY NOT NULL,
		consumer VARCHAR(20) NOT NULL,
		transformer VARCHAR(20) NOT NULL,
		address VARCHAR(30) NOT NULL,
		notification_type VARCHAR(20) DEFAULT "",
		deleted BOOLEAN DEFAULT 0,

		CONSTRAINT fk_contract_consumer FOREIGN KEY (consumer) REFERENCES consumer (id),
		CONSTRAINT fk_contract_transformer FOREIGN KEY (transformer) REFERENCES transformer (id)
	);
	
	CREATE TABLE reading
	(
		id VARCHAR(20) PRIMARY KEY NOT NULL,
		contract VARCHAR(20) NOT NULL,
		value INTEGER NOT NULL,
		date INTEGER NOT NULL,

		CONSTRAINT fk_reading_contract FOREIGN KEY (contract) REFERENCES contract (id)
	);

	CREATE TABLE bill
	(
		id VARCHAR(20) PRIMARY KEY NOT NULL,
		contract VARCHAR(20) NOT NULL,
		creation_date INTEGER NOT NULL,
		expiration_date INTEGER NOT NULL,
		paid BOOLEAN DEFAULT 0,

		CONSTRAINT fk_reading_contract FOREIGN KEY (contract) REFERENCES contract (id)
	);`

	_, err = DB.Exec(statement)
	if err != nil {
		return err
	}

	err = populate()
	if err != nil {
		return err
	}

	return nil
}
