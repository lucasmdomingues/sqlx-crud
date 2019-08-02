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

	db, err := newConn()
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	query := "INSERT INTO tb_users(id, username, password) VALUES(?, ?,?);"

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

func SelectRows() ([]User, error) {

	db, err := newConn()
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare("SELECT * FROM tb_users;")
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

		var (
			id                 int64
			username, password string
		)

		err = rows.Scan(&id, &username, &password)
		if err != nil {
			return nil, err
		}

		user := User{
			Id:       id,
			Username: username,
			Password: password,
		}

		users = append(users, user)
	}

	return users, nil
}

func SelectRow(id int) (*User, error) {

	db, err := newConn()
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare("SELECT * FROM tb_users WHERE id=?")
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

	db, err := newConn()
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	query := "UPDATE tb_users SET username=?, password=? WHERE id=?;"

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

	db, err := newConn()
	if err != nil {
		return nil
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	query := "DELETE FROM tb_users WHERE id=?"

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
