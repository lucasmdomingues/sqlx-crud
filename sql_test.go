package test

import (
	"go-crud-mysql/sql"
	"log"
	"testing"
)

func TestCreate(t *testing.T) {

	user := sql.User{
		1,
		"lucasmdomingues",
		"5490",
	}

	err := sql.Create(&user)
	if err != nil {
		t.Fatalf("Error: %v", err)
		return
	}
}

func TestSelectRows(t *testing.T) {

	users, err := sql.SelectRows()
	if err != nil {
		t.Fatalf("Error: %v", err)
		return
	}

	if len(users) == 0 {
		t.Fatal("No users found")
		return
	}

	log.Printf("%#v", users)
}

func TestUpdate(t *testing.T) {

	user := &sql.User{
		1,
		"lucasmdominguez",
		"5490",
	}

	err := sql.Update(user)
	if err != nil {
		t.Fatalf("Error: %v", err)
		return
	}
}

func TestSelectRow(t *testing.T) {

	user, err := sql.SelectRow(1)
	if err != nil {
		t.Fatalf("Error: %v", err)
		return
	}

	if user.Id == 0 {
		t.Fatalf("User not found")
		return
	}

	log.Printf("%#v", user)
}

func TestDelete(t *testing.T) {

	err := sql.Delete(1)
	if err != nil {
		t.Fatalf("Error: %v", err)
		return
	}
}
