package sql

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int64
	Username string
	Password string
}

func Create(user *User) error {

	query := "INSERT INTO tb_users(id, username, password) VALUES(?, ?, ?);"

	db, err := newConn()
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(&user.Id, &user.Username, &user.Password)
	if err != nil {
		tx.Rollback()
		return err
	}

	defer db.Close()
	tx.Commit()

	return nil
}

func FetchUsers() ([]User, error) {

	query := "SELECT * FROM tb_users;"

	db, err := newConn()
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	users := []User{}

	defer rows.Close()
	for rows.Next() {

		user := User{}

		err = rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func FetchUser(id int) (*User, error) {

	query := "SELECT * FROM tb_users WHERE id=?;"

	db, err := newConn()
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	user := User{}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
	}

	db.Close()

	return &user, nil
}

func Update(user *User) error {

	query := "UPDATE tb_users SET username=?, password=? WHERE id=?;"

	db, err := newConn()
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(&user.Username, &user.Password, &user.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	defer db.Close()
	tx.Commit()

	return nil
}

func Delete(id int) error {

	query := "DELETE FROM tb_users WHERE id=?;"

	db, err := newConn()
	if err != nil {
		return nil
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		return err
	}

	defer db.Close()
	tx.Commit()

	return nil
}
