package sql

import (
	"crypto/sha256"
	"fmt"
	"log"
	"testing"
)

func TestCreate(t *testing.T) {

	user := User{
		1,
		"lucasmdomingues",
		fmt.Sprintf("%x", sha256.Sum256([]byte("5490"))),
	}

	err := Create(&user)
	if err != nil {
		t.Fatalf("Error: %v", err)
		return
	}
}

func TestFetchUsers(t *testing.T) {

	users, err := FetchUsers()
	if err != nil {
		t.Fatalf("Error: %v", err)
		return
	}

	if len(users) == 0 {
		t.Fatal("No users found")
		return
	}

	log.Printf("%v", users)
}

func TestUpdate(t *testing.T) {

	user := &User{
		1,
		"lucasmdomingues",
		fmt.Sprintf("%x", sha256.Sum256([]byte("549054"))),
	}

	err := Update(user)
	if err != nil {
		t.Fatalf("Error: %v", err)
		return
	}
}

func TestFetchUser(t *testing.T) {

	user, err := FetchUser(1)
	if err != nil {
		t.Fatalf("Error: %v", err)
		return
	}

	if user.Id == 0 {
		t.Fatalf("User not found")
		return
	}

	log.Printf("%v", user)
}

func TestDelete(t *testing.T) {

	err := Delete(1)
	if err != nil {
		t.Fatalf("Error: %v", err)
		return
	}
}
