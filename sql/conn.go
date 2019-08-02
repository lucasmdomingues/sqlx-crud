package sql

import (
	"database/sql"
	"fmt"
)

const driverName = "mysql"

func newConn() (*sql.DB, error) {

	// User, Password, Host, Port, Database
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "root", "5490", "localhost", "3306", "db_users")

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
