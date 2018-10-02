package mysql

import (
	"database/sql"
	"fmt"
)

var urlConnection = makeURLConnetion()

type Connection struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func makeURLConnetion() string {

	c := Connection{
		"root",
		"",
		"localhost",
		"3306",
		"db_users",
	}

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.Username, c.Password, c.Host, c.Port, c.Database)

	return url
}

func makeConnection() *sql.DB {
	conn, err := sql.Open("mysql", urlConnection)
	if err != nil {
		return nil
	}

	return conn
}

var conn = makeConnection()

func closeConnetion(conn *sql.DB) {
	conn.Close()
}
