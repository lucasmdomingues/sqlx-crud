package mysql

import (
	"database/sql"
	"fmt"
	"log"
)

var conn *sql.DB = makeConnetion()

func makeConnetion() *sql.DB {

	username := ""
	password := ""
	host := ""
	port := ""
	database := ""

	connURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

	conn, err := sql.Open("mysql", connURL)
	if err != nil {
		log.Printf("%s", err)
	}
	return conn
}

func Insert() error {

	tx, err := conn.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Prepare("INSERT INTO tb_users(username,password) VALUES(?,?)")
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("user", "password")
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func Update(id int) error {

	tx, err := conn.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Prepare("update tb_users set username=?,password=? where uid=?")
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("newusername", "newpassoword", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func SelectAll() error {

	tx, err := conn.Begin()
	if err != nil {
		return err
	}

	rows, err := tx.Query("SELECT * FROM tb_users")
	if err != nil {
		tx.Rollback()
		return err
	}

	for rows.Next() {

		var id int
		var username string
		var password string

		err = rows.Scan(&id, &username, &password)
		if err != nil {
			tx.Rollback()
			return err
		}

		fmt.Println(id)
		fmt.Println(username)
		fmt.Println(password)
	}

	tx.Commit()
	return nil
}

func SelectWhere(id int) error {

	tx, err := conn.Begin()
	if err != nil {
		return err
	}

	rows, err := tx.Query("SELECT * tb_users WHERE id=?", id)
	if err != nil {
		return err
	}

	for rows.Next() {

		var id int
		var username string
		var password string

		err = rows.Scan(&id, &username, &password)
		if err != nil {
			tx.Rollback()
			return err
		}

		fmt.Println(id)
		fmt.Println(username)
		fmt.Println(password)
	}

	tx.Commit()
	return nil
}

func Delete(id int) error {
	stmt, err := conn.Prepare("DELETE FROM tb_users WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
