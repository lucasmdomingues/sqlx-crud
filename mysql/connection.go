package mysql

import (
	"database/sql"
	"fmt"
)

var conn = makeConnection()

type Connection struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func makeConnection() *sql.DB {

	c := Connection{
		"root",
		"",
		"localhost",
		"3306",
		"db_users",
	}

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.Username, c.Password, c.Host, c.Port, c.Database)

	conn, err := sql.Open("mysql", url)
	if err != nil {
		return nil
	}

	return conn
}

func closeConnetion(conn *sql.DB) {
	conn.Close()
}
