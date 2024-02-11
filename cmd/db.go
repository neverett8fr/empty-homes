package cmd

import (
	"database/sql"
	"empty-homes/pkg/config"
	"fmt"

	_ "github.com/lib/pq"
)

func OpenDB(conf *config.DB) (*sql.DB, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.Host,
		conf.Port,
		conf.Username,
		conf.Password,
		conf.Name,
	)

	conn, err := sql.Open(conf.Driver, connString)
	if err != nil {
		return nil, fmt.Errorf("error connecting to db, err %v", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("error sending ping to db, err %v", err)
	}

	return conn, nil
}
