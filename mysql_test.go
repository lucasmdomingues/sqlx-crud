package test

import (
	"golang-mysql/mysql"
	"testing"
)

func TestInsert(t *testing.T) {

	user := mysql.User{
		0,
		"lucasmdomingues",
		"5490",
	}

	err := mysql.Insert(&user)
	if err != nil {
		t.Fatalf("Error: %v", err)
		return
	}
}

func TestSelectAll(t *testing.T) {

	_, err := mysql.SelectAll()
	if err != nil {
		t.Fatalf("Error: %v", err)
		return
	}
}

func TestUpdate(t *testing.T) {

	user := mysql.User{
		0,
		"lucasmartinsd",
		"5490",
	}

	err := mysql.Update(&user, 0)
	if err != nil {
		t.Fatalf("Error: %v", err)
		return
	}
}

func TestSelectWhere(t *testing.T) {

	_, err := mysql.SelectWhere(0)
	if err != nil {
		t.Fatalf("Error: %v", err)
		return
	}
}

func TestDelete(t *testing.T) {

	err := mysql.Delete(0)
	if err != nil {
		t.Fatalf("Error: %v", err)
		return
	}
}
