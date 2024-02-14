package db

import (
	"fmt"
	"log"
	"time"
)

func (conn *DBConn) AddProperty(house Home) error {

	_, err := conn.Conn.Exec(
		"INSERT INTO properties(name, postcode, street, type, bedrooms, inhabited, safe, owner_name, owner_contact, date_added, date_last_checked) "+
			"VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		house.Name, house.Postcode, house.Street,
		house.Type, house.Bedrooms, house.Inhabited, house.Safe,
		house.OwnerName, house.OwnerContact,
		time.Now(), house.DateLastChecked,
	)
	if err != nil {
		err := fmt.Errorf("error adding property to db, err %v", err)
		log.Println(err)
		return err
	}

	return nil
}

func (conn *DBConn) ViewAll() ([]Home, error) {

	rows, err := conn.Conn.Query(fmt.Sprintf("SELECT id, name, postcode, street, type, bedrooms, inhabited, safe, owner_name, owner_contact, date_added, date_last_checked " +
		"FROM properties"))
	if err != nil {
		err := fmt.Errorf("error returning properties from db, err %v", err)
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var properites []Home
	for rows.Next() {
		var property Home
		if err := rows.Scan(&property.ID, &property.Name,
			&property.Postcode, &property.Street,
			&property.Type, &property.Bedrooms,
			&property.Inhabited, &property.Safe,
			&property.OwnerName, &property.OwnerContact,
			&property.DateAdded,
			&property.DateLastChecked); err != nil {
			return properites, err
		}
		properites = append(properites, property)

	}
	if err = rows.Err(); err != nil {
		return properites, err
	}

	return properites, nil
}
