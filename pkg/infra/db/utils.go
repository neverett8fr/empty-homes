package db

import "database/sql"

type DBConn struct {
	Conn *sql.DB
}

func NewDBConnFromExisting(conn *sql.DB) *DBConn {
	return &DBConn{
		Conn: conn,
	}
}
