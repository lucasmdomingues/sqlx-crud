package mysql

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int64
	Username string
	Password string
}

func Insert(user *User) error {

	conn := makeConnection()

	tx, err := conn.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	query := "INSERT INTO tb_users(username,password) VALUES(?,?)"

	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = stmt.Exec(&user.Username, &user.Password)
	if err != nil {
		return err
	}

	defer closeConnetion(conn)
	tx.Commit()

	return nil
}

func Update(user *User, idUser int64) error {

	conn := makeConnection()

	tx, err := conn.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	query := "UPDATE tb_users SET username=?,password=? WHERE id=?"

	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = stmt.Exec(&user.Username, &user.Password, idUser)
	if err != nil {
		return err
	}

	defer closeConnetion(conn)
	tx.Commit()

	return nil
}

func SelectAll() ([]User, error) {

	conn := makeConnection()

	tx, err := conn.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("SELECT * FROM tb_users")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	users := []User{}
	for rows.Next() {

		user := User{}

		err = rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		users = append(users, user)
	}

	defer closeConnetion(conn)
	tx.Commit()

	return users, nil
}

func SelectWhere(idUser int) (*User, error) {

	conn := makeConnection()

	tx, err := conn.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("SELECT * FROM tb_users WHERE id=?")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(idUser)
	if err != nil {
		return nil, err
	}

	user := User{}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	defer closeConnetion(conn)
	tx.Commit()

	return &user, nil
}

func Delete(idUser int) error {

	conn := makeConnection()

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

	defer closeConnetion(conn)
	tx.Commit()

	return nil
}
