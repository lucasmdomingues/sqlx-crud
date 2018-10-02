package test

import (
	"golang-mysql/mysql"
	"log"
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
		t.Errorf("Error: %v", err)
		return
	}

	log.Printf("Test func Insert return success\n")
}

func TestSelectAll(t *testing.T) {

	user, err := mysql.SelectAll()
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}

	log.Printf("Test func SelectAll return: %v\n", user)
}

func TestUpdate(t *testing.T) {

	user := mysql.User{
		0,
		"lucasmartinsd",
		"5490",
	}

	err := mysql.Update(&user, 0)
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}

	log.Printf("Test func Insert return success\n")
}

func TestSelectWhere(t *testing.T) {

	user, err := mysql.SelectWhere(0)
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}

	log.Printf("Test func SelectWhere return: %v\n", user)
}

func TestDelete(t *testing.T) {

	err := mysql.Delete(0)
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}

	log.Printf("Test func Delete return success")

}
