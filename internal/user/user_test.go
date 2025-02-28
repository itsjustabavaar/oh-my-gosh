package user

import (
	_ "github.com/itsjustabavaar/oh-my-gosh/internal/database"
	"testing"
	_ "time"
)

func TestUserScenario(t *testing.T) {
	username, password := "testuser2", "1234"

	err := AddUser(username, password)
	if err != nil {
		t.Fatalf("Error adding user: %v", err)
	}

	user, err := Login(username, password)
	if err != nil {
		t.Fatalf("Error logging in: %v", err)
	}

	if user.Username != username {
		t.Fatalf("Username does not match")
	}

	err = DeleteUser(username)
	if err != nil {
		t.Fatalf("Error deleting user: %v", err)
	}
}

func TestUserNotFoundLogin(t *testing.T) {
	username, password := "testuser2", "1234"

	err := AddUser(username, password)
	if err != nil {
		t.Fatalf("Error adding user: %v", err)
	}

	err = DeleteUser(username)
	if err != nil {
		t.Fatalf("Error deleting user: %v", err)
	}

	_, err = Login(username, password)
	if err != nil {
		if err.Error() != "user not found" {
			t.Fatal("unexpected error")
		}
	}
}

func TestUserAlreadyExists(t *testing.T) {
	username, password := "testuser2", "1234"
	err := AddUser(username, password)
	if err != nil {
		t.Fatalf("Error adding user: %v", err)
	}

	err = AddUser(username, password)
	if err != nil {
		if err.Error() != "user already exists" {
			t.Fatal("unexpected error")
		}
	}
	err = DeleteUser(username)
	if err != nil {
		t.Fatalf("Error deleting user: %v", err)
	}
}

func TestLoginInvalidPassword(t *testing.T) {
	username, password := "testuser2", "1234"
	err := AddUser(username, password)
	if err != nil {
		t.Fatalf("Error adding user: %v", err)
	}

	_, err = Login(username, "1235")
	if err != nil {
		if err.Error() != "invalid password" {
			t.Fatal("unexpected error")
		}
	}
	err = DeleteUser(username)
	if err != nil {
		t.Fatalf("Error deleting user: %v", err)
	}
}
