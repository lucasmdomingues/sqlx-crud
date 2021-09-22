package user

import "lucasmdomingues/gosql-crud/database"

func Create(user *User) error {
	db, err := database.NewConnection()
	if err != nil {
		return err
	}

	tx := db.MustBegin()

	_, err = tx.NamedExec("INSERT INTO user (username, password) VALUES (:username, :password)", &user)
	if err != nil {
		tx.Rollback()

		return err
	}
	defer db.Close()

	tx.Commit()

	return nil
}

func FindUsers() ([]User, error) {
	db, err := database.NewConnection()
	if err != nil {
		return nil, err
	}

	users := make([]User, 0)

	if err = db.Select(&users, "SELECT * FROM user"); err != nil {
		return nil, err
	}
	defer db.Close()

	return users, nil
}

func GetByID(id int) (*User, error) {
	db, err := database.NewConnection()
	if err != nil {
		return nil, err
	}

	var user User

	if err = db.Get(&user, "SELECT * FROM user WHERE id = ?", id); err != nil {
		return nil, err
	}
	defer db.Close()

	return &user, nil
}

func Update(user *User) error {
	db, err := database.NewConnection()
	if err != nil {
		return err
	}

	tx := db.MustBegin()

	_, err = tx.NamedExec("UPDATE user SET username=:username, password=:password WHERE id=:id", &user)
	if err != nil {
		tx.Rollback()

		return err
	}
	defer db.Close()

	tx.Commit()

	return nil
}

func Delete(id int) error {
	db, err := database.NewConnection()
	if err != nil {
		return err
	}

	tx := db.MustBegin()

	tx.MustExec("DELETE FROM user WHERE id=?", id)
	defer db.Close()

	tx.Commit()

	return nil
}
