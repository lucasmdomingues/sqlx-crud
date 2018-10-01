package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var conn *sql.DB = makeConnetion()

type User struct {
	Id       int64
	Username string
	Password string
}

func makeConnetion() *sql.DB {

	username := "root"
	password := ""
	host := "localhost"
	port := "3306"
	database := "orcamentoja.kouda_db"

	connURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

	conn, err := sql.Open("mysql", connURL)
	if err != nil {
		log.Printf("%s", err)
	}
	return conn
}

func Insert(user User) error {

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

	_, err = tx.Exec(user.Username, user.Password)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func Update(user User, idUser int) error {

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

	_, err = tx.Exec(user.Username, user.Password, idUser)
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
	return nil
}

func Delete(idUser int) error {
	stmt, err := conn.Prepare("DELETE FROM tb_users WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(idUser)
	if err != nil {
		return err
	}

	return nil
}
