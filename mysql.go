package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var conn *sql.DB = makeConnetion()

type Connection struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

type User struct {
	Id       int64
	Username string
	Password string
}

func makeConnetion() *sql.DB {

	c := Connection{
		"username",
		"password",
		"host",
		"port",
		"database",
	}

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.Username, c.Password, c.Host, c.Port, c.Database)

	conn, err := sql.Open("mysql", url)
	if err != nil {
		return nil
	}

	return conn
}

func Insert(user User) error {

	tx, err := conn.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	query := "INSERT INTO tb_users(username,password) VALUES(?,?)"

	_, err = tx.Exec(query, user.Username, user.Password)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	conn.Close()

	return nil
}

func Update(user User, idUser int) error {

	tx, err := conn.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	query := "update tb_users set username=?,password=? where uid=?"

	_, err = tx.Exec(query, user.Username, user.Password, idUser)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	conn.Close()

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

		var user User

		err = rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			tx.Rollback()
			return err
		}

		fmt.Println(user.Id)
		fmt.Println(user.Username)
		fmt.Println(user.Password)
	}

	tx.Commit()
	conn.Close()

	return nil
}

func SelectWhere(idUser int) error {

	tx, err := conn.Begin()
	if err != nil {
		return err
	}

	rows, err := tx.Query("SELECT * tb_users WHERE id=?", idUser)
	if err != nil {
		return err
	}

	for rows.Next() {

		var user User

		err = rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			tx.Rollback()
			return err
		}

		fmt.Println(user.Id)
		fmt.Println(user.Username)
		fmt.Println(user.Password)
	}

	tx.Commit()
	conn.Close()

	return nil
}

func Delete(idUser int) error {

	tx, err := conn.Begin()
	if err != nil {
		return err
	}

	query := "DELETE FROM tb_users WHERE id=?"

	_, err = tx.Exec(query, idUser)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	conn.Close()

	return nil
}
