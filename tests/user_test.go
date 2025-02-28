package tests

import (
	_ "github.com/itsjustabavaar/oh-my-gosh/internal/database"
	user2 "github.com/itsjustabavaar/oh-my-gosh/internal/user"
	"testing"
	_ "time"
)

func TestUserScenario(t *testing.T) {
	username, password := "testuser2", "1234"

	err := user2.AddUser(username, password)
	if err != nil {
		t.Fatalf("Error adding user: %v", err)
	}

	user, err := user2.Login(username, password)
	if err != nil {
		t.Fatalf("Error logging in: %v", err)
	}

	if user.Username != username {
		t.Fatalf("Username does not match")
	}

	err = user2.DeleteUser(username)
	if err != nil {
		t.Fatalf("Error deleting user: %v", err)
	}
}

func TestUserNotFoundLogin(t *testing.T) {
	username, password := "testuser2", "1234"

	err := user2.AddUser(username, password)
	if err != nil {
		t.Fatalf("Error adding user: %v", err)
	}

	err = user2.DeleteUser(username)
	if err != nil {
		t.Fatalf("Error deleting user: %v", err)
	}

	_, err = user2.Login(username, password)
	if err != nil {
		if err.Error() != "user not found" {
			t.Fatal("unexpected error")
		}
	}
}

func TestUserAlreadyExists(t *testing.T) {
	username, password := "testuser2", "1234"
	err := user2.AddUser(username, password)
	if err != nil {
		t.Fatalf("Error adding user: %v", err)
	}

	err = user2.AddUser(username, password)
	if err != nil {
		if err.Error() != "user already exists" {
			t.Fatal("unexpected error")
		}
	}
	err = user2.DeleteUser(username)
	if err != nil {
		t.Fatalf("Error deleting user: %v", err)
	}
}

func TestUserLoginInvalidPassword(t *testing.T) {
	username, password := "testuser2", "1234"
	err := user2.AddUser(username, password)
	if err != nil {
		t.Fatalf("Error adding user: %v", err)
	}

	_, err = user2.Login(username, "1235")
	if err != nil {
		if err.Error() != "invalid password" {
			t.Fatal("unexpected error")
		}
	}
	err = user2.DeleteUser(username)
	if err != nil {
		t.Fatalf("Error deleting user: %v", err)
	}
}
