package storage

import (
	"fmt"
)

func populate() error {
	statement := `
	INSERT INTO employee(email, password, type, deleted) VALUES("admin@employee.com", "123", "admin", false);
	INSERT INTO employee(email, password, type, deleted) VALUES("operator@employee.com", "123", "operator", false);
	INSERT INTO employee(email, password, type, deleted) VALUES("manager@employee.com", "123", "manager", false);
	INSERT INTO substation(id, name, longitude, latitude, deleted) VALUES("SUB1234567890123456", "Univalle", "-76.53337259999999", "3.3758083", false);
	INSERT INTO substation(id, name, longitude, latitude, deleted) VALUES("SUB1234567890123457", "Pailados", "-76.531294", "3.383681", false);
	INSERT INTO transformer(id, name, substation, longitude, latitude, deleted) VALUES("TRA1234567890123456", "Pance", "SUB1234567890123456", "-76.63865", "3.32834", false);
	INSERT INTO transformer(id, name, substation, longitude, latitude, deleted) VALUES("TRA1234567890123457", "Sonido", "SUB1234567890123456", "-76.5323217", "3.4527262", false);
	INSERT INTO transformer(id, name, substation, longitude, latitude, deleted) VALUES("TRA1234567890123458", "Truora", "SUB1234567890123457", "-76.5321768", "3.3831216", false);
	INSERT INTO transformer(id, name, substation, longitude, latitude, deleted) VALUES("TRA1234567890123459", "747", "SUB1234567890123457", "-76.53301139999999", "3.3822217", false);
	INSERT INTO consumer(id, password, email, deleted) VALUES("1111111", "123", "shadyelectriccompany@gmail.com", false);
	INSERT INTO consumer(id, password, email, deleted) VALUES("2222222", "123", "shadyelectriccompany@gmail.com", false);
	INSERT INTO contract(id, consumer, transformer, address, notification_type, deleted) VALUES("CON1234567890123456", "1111111", "TRA1234567890123456", "address 1", "email", false);
	INSERT INTO contract(id, consumer, transformer, address, notification_type, deleted) VALUES("CON1234567890123457", "1111111", "TRA1234567890123456", "address 2", "email", false);
	INSERT INTO contract(id, consumer, transformer, address, notification_type, deleted) VALUES("CON1234567890123458", "2222222", "TRA1234567890123456", "address 3", "email", false);`

	// ////////////////////////////////////////
	// ////////////////////////////////////////
	// Contract CON1234567890123456
	// Last month
	// Readings of month 4
	statement += `INSERT INTO reading(id, contract, value, date) VALUES("REA000000000000000", "CON1234567890123456", 100, 1586908800);`

	// ////////////////////////////////////////
	// ////////////////////////////////////////
	// Contract CON1234567890123457
	// Last month
	// Readings of month 4
	statement += `INSERT INTO reading(id, contract, value, date) VALUES("REA1111111111111111", "CON1234567890123457", 120, 1586908800);`

	// Last last month (pending pay)
	// Readings of month 3
	// Bill of month 4
	statement += `INSERT INTO reading(id, contract, value, date) VALUES("REA1111111111111112", "CON1234567890123457", 120, 1584230400);`
	statement += `INSERT INTO bill(id, contract, creation_date, expiration_date, paid) VALUES("BIL1111111111111111", "CON1234567890123457", 1585699200, 1588204800, false);`

	// Last 5 paid months
	// Readings of month 2
	// Bill of month 3
	statement += `INSERT INTO reading(id, contract, value, date) VALUES("REA1111111111111113", "CON1234567890123457", 120, 1581724800);`
	statement += `INSERT INTO bill(id, contract, creation_date, expiration_date, paid) VALUES("BIL1111111111111112", "CON1234567890123457", 1583020800, 1585612800, true);`

	// Readings of month 1
	// Bill of month 2
	statement += `INSERT INTO reading(id, contract, value, date) VALUES("REA1111111111111114", "CON1234567890123457", 120, 1579046400);`
	statement += `INSERT INTO bill(id, contract, creation_date, expiration_date, paid) VALUES("BIL1111111111111113", "CON1234567890123457", 1580515200, 1582934400, true);`

	// Readings of month 12
	// Bill of month 1
	statement += `INSERT INTO reading(id, contract, value, date) VALUES("REA1111111111111115", "CON1234567890123457", 120, 1576368000);`
	statement += `INSERT INTO bill(id, contract, creation_date, expiration_date, paid) VALUES("BIL1111111111111114", "CON1234567890123457", 1577836800, 1580428800, true);`

	// Readings of month 11
	// Bill of month 12
	statement += `INSERT INTO reading(id, contract, value, date) VALUES("REA1111111111111116", "CON1234567890123457", 120, 1573776000);`
	statement += `INSERT INTO bill(id, contract, creation_date, expiration_date, paid) VALUES("BIL1111111111111115", "CON1234567890123457", 1575158400, 1577750400, true);`

	// Readings of month 10
	// Bill of month 11
	statement += `INSERT INTO reading(id, contract, value, date) VALUES("REA1111111111111117", "CON1234567890123457", 120, 1571097600);`
	statement += `INSERT INTO bill(id, contract, creation_date, expiration_date, paid) VALUES("BIL1111111111111116", "CON1234567890123458", 1572566400, 1575072000, true);`

	// ////////////////////////////////////////
	// ////////////////////////////////////////
	// Contract CON1234567890123458
	// Last month
	// Readings of month 4
	statement += `INSERT INTO reading(id, contract, value, date) VALUES("REA2222222222222221", "CON1234567890123458", 120, 1586908800);`

	// Last last month (pending pay)
	// Readings of month 3
	// Bill of month 4
	statement += `INSERT INTO reading(id, contract, value, date) VALUES("REA2222222222222222", "CON1234567890123458", 120, 1584230400);`
	statement += `INSERT INTO bill(id, contract, creation_date, expiration_date, paid) VALUES("BIL2222222222222221", "CON1234567890123458", 1585699200, 1588204800, false);`

	// Last last last month (pending pay)
	// Readings of month 2
	// Bill of month 3
	statement += `INSERT INTO reading(id, contract, value, date) VALUES("REA2222222222222223", "CON1234567890123458", 120, 1581724800);`
	statement += `INSERT INTO bill(id, contract, creation_date, expiration_date, paid) VALUES("BIL2222222222222222", "CON1234567890123458", 1583020800, 1585612800, false);`

	// Last 3 paid months
	// Readings of month 1
	// Bill of month 2
	statement += `INSERT INTO reading(id, contract, value, date) VALUES("REA2222222222222224", "CON1234567890123458", 120, 1579046400);`
	statement += `INSERT INTO bill(id, contract, creation_date, expiration_date, paid) VALUES("BIL2222222222222223", "CON1234567890123458", 1580515200, 1582934400, true);`

	// Readings of month 12
	// Bill of month 1
	statement += `INSERT INTO reading(id, contract, value, date) VALUES("REA2222222222222225", "CON1234567890123458", 120, 1576368000);`
	statement += `INSERT INTO bill(id, contract, creation_date, expiration_date, paid) VALUES("BIL2222222222222224", "CON1234567890123458", 1577836800, 1580428800, true);`

	// Readings of month 11
	// Bill of month 12
	statement += `INSERT INTO reading(id, contract, value, date) VALUES("REA2222222222222226", "CON1234567890123458", 120, 1573776000);`
	statement += `INSERT INTO bill(id, contract, creation_date, expiration_date, paid) VALUES("BIL2222222222222225", "CON1234567890123458", 1575158400, 1577750400, true);`

	result, err := DB.Exec(statement)
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
